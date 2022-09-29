var video = document.getElementById('video');
var sources = document.getElementById('source');
var isLoad = false
setInterval(loop, 1000);

video.onplay = function () {
    showVideo();
    isLoad = true;
}

video.onended = function () {
    loadVideo();
}

sources.onerror = function () {
    hideVideo();
    isLoad = false
    console.clear();
}

function loop() {
    if (!isLoad) {
        loadVideo();
    }
}

function loadVideo() {
    sources.src = "/content?time=" + new Date().getTime();
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