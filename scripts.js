function refresh() {
    let promise = fetch("http://localhost:8090/refresh/")
}
window.addEventListener("load", refresh) // call refresh on opening the timer

function startTimer(buttonNumber) {
    fetch("http://localhost:8090/start/" + buttonNumber, {mode:'no-cors'});
    console.log("Starting timer...");
    }

function stopTimer(buttonNumber) {
    fetch("http://localhost:8090/stop/" + buttonNumber, {mode:'no-cors'});
    console.log("Stopping timer...");
    }

function getRunning(buttonNumber) {
    fetch("http://localhost:8090/getrunning/" + buttonNumber).then
        (response => response.json())
        .then(data => console.log(data));
}