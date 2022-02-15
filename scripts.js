function refresh() {
  let promise = fetch("http://localhost:8090/update/");
}
window.addEventListener("load", refresh); // call refresh on opening the timer

async function startTimer(buttonNumber) {
  fetch("http://localhost:8090/start/" + buttonNumber, { mode: "no-cors" });
  console.log("Starting timer...");
}

async function stopTimer(buttonNumber) {
  fetch("http://localhost:8090/stop/" + buttonNumber, { mode: "no-cors" });
  console.log("Stopping timer...");
}

async function getRunning(buttonNumber) {
  fetch("http://localhost:8090/getrunning/" + buttonNumber)
    .then((response) => response.json())
    .then((data) => console.log(data));
}

const timer1 = document.getElementById("timer1");
const timer2 = document.getElementById("timer2");
const timer3 = document.getElementById("timer3");
const timer4 = document.getElementById("timer4");
const timer5 = document.getElementById("timer5");
const timer6 = document.getElementById("timer6");
const timer7 = document.getElementById("timer7");
const timer8 = document.getElementById("timer8");
const timer9 = document.getElementById("timer9");
const timer10 = document.getElementById("timer10");
const timer11 = document.getElementById("timer11");
const timer12 = document.getElementById("timer12");
const timer13 = document.getElementById("timer13");
const timer14 = document.getElementById("timer14");
const timer15 = document.getElementById("timer15");
const timer16 = document.getElementById("timer16");
const timer17 = document.getElementById("timer17");
const timer18 = document.getElementById("timer18");
const timer19 = document.getElementById("timer19");
const timer20 = document.getElementById("timer20");

function delay(time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}

window.addEventListener("load", Update);

let updateInit = false;

async function Update() {
  // Ensure Update is run only once
  if (updateInit) {
    // Update already initialized
    return;
  }

  updateInit = true;

  console.log("Begin update");
  while (updateInit) {
    for (let i = 1; i < 21; i++) {
      console.log("Fetching state");
      fetch("http://localhost:8090/getstate/" + i)
      .then((response) => response.json())
      .then((data) => RedrawTimers(data.Id, data.Elapsed));
    //   .then((data) => RedrawTimers(data.Id, data.Elapsed, data.Running, data.OutOfOrder));
    }
    await(delay(1000));
  }
}

async function RedrawTimers(id, elapsed) {
    let target = document.getElementById("timer" + id)
    target.innerHTML = elapsed
}
