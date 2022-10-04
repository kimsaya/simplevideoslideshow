var video = document.getElementById('video');
var image = document.getElementById('image');
var sources = document.getElementById('source');
var duration = 1000;
var mainLoop = setInterval(loop, duration);

var getJSON = function (url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function () {
        var status = xhr.status;
        if (status === 200) {
            callback(null, xhr.response);
        } else {
            callback(status, xhr.response);
        }
    };
    xhr.send();
};

function getContent() {
    getJSON('/content',
        function (err, data) {
            if (err !== null) {
                alert('Read Content Failed ' + err);
            } else {
                // alert('Your query count: ' + data.name);
                if (data.item_type == "video") {
                    sources.src = "/play?url=" + data.url;
                    video.load();
                }
            }
        });
}

// Video function
video.onplay = function () {
    showVideo();
    stopLoop();
}

video.onended = function () {
    loadVideo();
}

sources.onerror = function () {
    hideVideo();
    setLoop();
    console.clear();
}

function setLoop() {
    stopLoop();
    mainLoop = setInterval(loop, duration);
}

function stopLoop() {
    clearInterval(mainLoop);
}

function loop() {
    // loadVideo();
    getContent();
}

function loadVideo() {
    sources.src = "/nextvideo?time=" + new Date().getTime();
    video.load();
}

function hideVideo() {
    if (video.style.display === "none") {
        return;
    }
    video.style.display = "none";

}

function showVideo() {
    if (video.style.display === "block") {
        return;
    }
    video.style.display = "block";
}

function hideImage() {
    if (image.style.display === "none") {
        return;
    }
    image.style.display = "none";

}

function showImage() {
    if (image.style.display === "block") {
        return;
    }
    image.style.display = "block";
}