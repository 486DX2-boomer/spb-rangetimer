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

async function clearTimer(buttonNumber) {
  fetch("http://localhost:8090/clear/" + buttonNumber, { mode: "no-cors" });
  console.log("Stopping timer...");
}

async function getRunning(buttonNumber) {
  fetch("http://localhost:8090/getrunning/" + buttonNumber)
    .then((response) => response.json())
    .then((data) => console.log(data));
}
async function setOutOfOrder(buttonNumber) {
  fetch("http://localhost:8090/outoforder/" + buttonNumber, { mode: "no-cors" });
}

async function setMember(buttonNumber) {
  fetch("http://localhost:8090/member/" + buttonNumber, { mode: "no-cors" });

}

async function setReserved(buttonNumber) {
  fetch("http://localhost:8090/reserved/" + buttonNumber, { mode: "no-cors" });
}

function delay(time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}
// Convert an integer to a string in the format 00:00:00 (eg 3600 to 01:00:00)
function FormatTime(seconds) {
  let time = new Date(1000 * seconds).toISOString().substr(11, 8);
  return time;
}

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
        .then((data) => RedrawTimers(data.Id, data.Elapsed, data.Member, data.Reserved, data.OutOfOrder, data.Expired));
    }
    await delay(1000);
  }
}
window.addEventListener("load", Update);

async function RedrawTimers(id, elapsed, member, reserved, outOfOrder, expired) {
  let target = document.getElementById("timer" + id);
  let tn = document.getElementById("tn" + id);
  target.innerHTML = FormatTime(elapsed);
  if (member) {
    tn.innerHTML = (id + " - " + "Member")
  } else if (!member) {
    tn.innerHTML = (id)
  }

  if (reserved) {
    target.innerHTML = "Reserved"
  }

  if (outOfOrder) {
    target.innerHTML = `<span class="outOfOrderMessage">Out of Order</span>`
  }

  if (expired) {
    target.innerHTML = `<span class="expiredMessage">00:00:00</span>`
  }

}

hotkeys("ctrl+shift+h", function (event, handler) {
  switch (handler.key) {
    case "ctrl+shift+h":
      hideButtons();
      break;
    default:
      alert(event);
  }
});

let buttonsHidden = false;

function hideButtons() {
  let p = document.getElementsByClassName("button-container");

  if (buttonsHidden == false) {
    for (i = 0; i < p.length; i++) {
      p[i].style.display = "none";
    }
    
    console.log("Buttons hidden");
    buttonsHidden = true;
    
  } else if (buttonsHidden == true) {
    for (i = 0; i < p.length; i++) {
      p[i].style.display = "block";
    }
    
    console.log("Buttons visible");
    buttonsHidden = false;

  }
}
