$(function () {
    var FADE_TIME = 150; // ms
    var TYPING_TIMER_LENGTH = 400; // ms
    var COLORS = [
        '#e21400', '#91580f', '#f8a700', '#f78b00',
        '#58dc00', '#287b00', '#a8f07a', '#4ae8c4',
        '#3b88eb', '#3824aa', '#a700ff', '#d300e7'
    ];

    // Initialize variables
    var $window = $(window);
    var $usernameInput = $('.usernameInput'); // Input for username
    var $messages = $('.messages'); // Messages area
    var $inputMessage = $('.inputMessage'); // Input message input box

    var $loginPage = $('.login.page'); // The login page
    var $chatPage = $('.chat.page'); // The chatroom page

    // Prompt for setting a username
    var username;
    var connected = false;
    var typing = false;
    var lastTypingTime;
    var $currentInput = $usernameInput.focus();

    var socket = io();

    const addParticipantsMessage = (data) => {
        var message = '';

        if (data.numUsers === 1) {
            // message += "there's 1 participant";
            message += "现在好像就你一个人";
        } else {
            message += "有 " + data.numUsers + " 个人了";
        }
        log(message);
    }

    // Sets the client's username
    const setUsername = () => {
        username = cleanInput($usernameInput.val().trim());

        // If the username is valid
        if (username) {
            $loginPage.fadeOut();
            $chatPage.show();
            $loginPage.off('click');
            $currentInput = $inputMessage.focus();

            // Tell the server your username
            socket.emit('add user', username);
        }
    }

    // Sends a chat message
    const sendMessage = () => {
        var message = $inputMessage.val();
        // Prevent markup from being injected into the message
        message = cleanInput(message);
        // if there is a non-empty message and a socket connection
        if (message && connected) {
            $inputMessage.val('');
            addChatMessage({
                username: username,
                message: {
                    type: 0,
                    message: message
                }
            });
            var messages = {'type': 0, 'message': message}
            // tell server to execute 'new message' and send along one parameter
            socket.emit('new message', messages);
        }
    }

    // ------ 发送图片 -------
    var file = document.getElementById('file')
    file.onchange = function () {
        imgf = file.files[0]
        lrz(imgf).then(function (res) {
            console.log('压缩图片结果：' + res.base64)
            var reader = new FileReader()
            reader.readAsDataURL(imgf)
            reader.onload = function () {
                addChatMessage({
                    username: username,
                    message: {
                        type: 1,
                        message: res.base64
                    }
                });
            }
            upload_img(res.file, res.origin.name)
        })
    }

    function upload_img(file, filename) {
        var csrfToken = $('meta[name=csrf-token]').attr('content')
        var formdata = new FormData()
        formdata.append('images', file)
        formdata.append('filename', filename)
        $.ajax({
            url: '/files/upload',
            type: "post",
            data: formdata,
            contentType: false,
            processData: false,
            beforeSend: function (xhr,settings) {
                xhr.setRequestHeader("X-CSRFToken",csrfToken)
                // $("#loading").html("<img src='./images/loading.gif' />");
            },
            success: function (data) {
                console.log(data.data)
                // var datas = JSON.parse(data)
                var messages = {'type': 1, 'message': data.data}
                // tell server to execute 'new message' and send along one parameter
                socket.emit('new message', messages);
                // alert('上传成功'+data)
            },
            error: function (data) {
                console.log(data.toString())
                alert('上传失败' + data)
            }

        });
    }

    //^^^^^^^^^^^^^^^^^
    $('#img_box').hide()
    // Log a message
    const log = (message, options) => {
        var $el = $('<li>').addClass('log').text(message);
        addMessageElement($el, options);
    }

    // Adds the visual chat message to the message list
    const addChatMessage = (data, options) => {
        var username = data.username
        var type = data.message.type
        var message = data.message.message
        console.log(username, message)
        // Don't fade the message in if there is an 'X was typing'
        var $typingMessages = getTypingMessages(data);
        options = options || {};
        if ($typingMessages.length !== 0) {
            options.fade = false;
            $typingMessages.remove();
        }

        var $usernameDiv = $('<span class="username"/>')
            .text(username)
            .css('color', getUsernameColor(username));
        var $messageBodyDiv = $('<span class="messageBody">')
            .text(message);

        var typingClass = data.typing ? 'typing' : '';
        if (type === 0) {//文本
            var $messageDiv = $('<li class="message"/>')
                .data('username', username)
                .addClass(typingClass)
                .append($usernameDiv, $messageBodyDiv);
        } else if (type === 1) { //图片
            // var abc = 'http://192.168.31.54:5000/files/image/确认信息-行驶证.png'
            var img = $("<img class='img' >").attr('src', message).css({
                'width': '100px',
                'height': '100px',
                'object-fit': 'cover'
            }).click(function (event) {
                $('#img_box').show()
                $('#img_origin').attr('src',this.src)
            })
            var $messageDiv = $('<li class="message"/>')
                .data('username', username)
                .addClass(typingClass)
                .append($usernameDiv, img);
        }
        addMessageElement($messageDiv, options);
    }

    // Adds the visual chat typing message
    const addChatTyping = (data) => {
        data.typing = true;
        data.message = 'is typing';
        addChatMessage(data);
    }

    // Removes the visual chat typing message
    const removeChatTyping = (data) => {
        getTypingMessages(data).fadeOut(function () {
            $(this).remove();
        });
    }

    // Adds a message element to the messages and scrolls to the bottom
    // el - The element to add as a message
    // options.fade - If the element should fade-in (default = true)
    // options.prepend - If the element should prepend
    //   all other messages (default = false)
    const addMessageElement = (el, options) => {
        var $el = $(el);

        // Setup default options
        if (!options) {
            options = {};
        }
        if (typeof options.fade === 'undefined') {
            options.fade = true;
        }
        if (typeof options.prepend === 'undefined') {
            options.prepend = false;
        }

        // Apply options
        if (options.fade) {
            $el.hide().fadeIn(FADE_TIME);
        }
        if (options.prepend) {
            $messages.prepend($el);
        } else {
            $messages.append($el);
        }
        $messages[0].scrollTop = $messages[0].scrollHeight;
    }

    // Prevents input from having injected markup
    const cleanInput = (input) => {
        return $('<div/>').text(input).html();
    }

    // Updates the typing event
    const updateTyping = () => {
        if (connected) {
            if (!typing) {
                typing = true;
                socket.emit('typing');
            }
            lastTypingTime = (new Date()).getTime();

            setTimeout(() => {
                var typingTimer = (new Date()).getTime();
                var timeDiff = typingTimer - lastTypingTime;
                if (timeDiff >= TYPING_TIMER_LENGTH && typing) {
                    socket.emit('stop typing');
                    typing = false;
                }
            }, TYPING_TIMER_LENGTH);
        }
    }

    // Gets the 'X is typing' messages of a user
    const getTypingMessages = (data) => {
        return $('.typing.message').filter(function (i) {
            return $(this).data('username') === data.username;
        });
    }

    // Gets the color of a username through our hash function
    const getUsernameColor = (username) => {
        // Compute hash code
        var hash = 7;
        for (var i = 0; i < username.length; i++) {
            hash = username.charCodeAt(i) + (hash << 5) - hash;
        }
        // Calculate color
        var index = Math.abs(hash % COLORS.length);
        return COLORS[index];
    }

    // Keyboard events

    $window.keydown(event => {
        // Auto-focus the current input when a key is typed
        if (!(event.ctrlKey || event.metaKey || event.altKey)) {
            $currentInput.focus();
        }
        // When the client hits ENTER on their keyboard
        if (event.which === 13) {
            if (username) {
                sendMessage();
                socket.emit('stop typing');
                typing = false;
            } else {
                setUsername();
            }
        }
    });

    $inputMessage.on('input', () => {
        updateTyping();
    });

    // Click events

    // Focus input when clicking anywhere on login page
    $loginPage.click(() => {
        $currentInput.focus();
    });

    // Focus input when clicking on the message input's border
    $inputMessage.click(() => {
        $inputMessage.focus();
    });

    // Socket events

    // Whenever the server emits 'login', log the login message
    socket.on('login', (data) => {
        console.log(data)
        connected = true;
        // Display the welcome message
        var message = "来了就说句话吧 – ";
        log(message, {
            prepend: true
        });
        addParticipantsMessage(data);
    });

    // Whenever the server emits 'new message', update the chat body
    socket.on('new message', (data) => {
        addChatMessage(data);
    });

    // Whenever the server emits 'user joined', log it in the chat body
    socket.on('user joined', (data) => {
        log(data.username + ' joined');
        addParticipantsMessage(data);
    });

    // Whenever the server emits 'user left', log it in the chat body
    socket.on('user left', (data) => {
        log(data.username + ' left');
        addParticipantsMessage(data);
        removeChatTyping(data);
    });

    // Whenever the server emits 'typing', show the typing message
    socket.on('typing', (data) => {
        addChatTyping(data);
    });

    // Whenever the server emits 'stop typing', kill the typing message
    socket.on('stop typing', (data) => {
        removeChatTyping(data);
    });

    socket.on('disconnect', () => {
        log('you have been disconnected');
    });

    socket.on('reconnect', () => {
        log('you have been reconnected');
        if (username) {
            socket.emit('add user', username);
        }
    });

    socket.on('reconnect_error', () => {
        log('attempt to reconnect has failed');
    });

});