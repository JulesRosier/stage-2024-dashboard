window.onload = (e) => {
  const formTemplate = document.getElementById("form-template");
  const queries = document.getElementById("queries");
  const addFormBtn = document.getElementById("add-form");
  const nerdInput = document.querySelector('[name="nerd_mode"]');
  addFormBtn.onclick = () => {
    queries.appendChild(formTemplate.content.cloneNode(true));
    addL();
    setHeight();
  };
  Sortable.create(queries, { animation: 150, onUpdate: reload });
  addL();
  document.cookie =
    "timezone=" + Intl.DateTimeFormat().resolvedOptions().timeZone;

  nerdInput.checked = localStorage.getItem("nerd_mode") == "true";
  nerdInput.onchange = (event) => {
    localStorage.setItem("nerd_mode", event.target.checked);
  };
};

function addL() {
  let removeBtns;
  removeBtns = document.getElementsByClassName("rm-form");
  for (let b of removeBtns) {
    b.onclick = (e) => {
      e.target.parentElement.parentElement.remove();
      reload();
      setHeight();
    };
  }
}

function addQuery(index, query) {
  const formTemplate = document.getElementById("form-template");
  const queries = document.getElementById("queries");
  const q = formTemplate.content.cloneNode(true);
  const select = q.firstChild.firstChild.children[0];
  const input = q.firstChild.firstChild.children[1];
  const btn = q.firstChild.firstChild.children[2];
  btn.onclick = (e) => {
    e.target.parentElement.parentElement.remove();
    reload();
  };
  select.value = index;
  input.value = query;
  queries.appendChild(q);
  reload();
}

function reload() {
  htmx.trigger("#main-form", "onLoadC");
}

function hover() {
  const hoverElements = document.querySelectorAll('[class^="hhh-"]');
  hoverElements.forEach((element) => {
    element.addEventListener("mouseover", () => {
      const subclass = element.classList[0]; // Get the class of the hovered element
      const elementsInSubclass = document.querySelectorAll(`.${subclass}`);
      elementsInSubclass.forEach((el) => {
        el.classList.add("hover-color");
      });
    });
    element.addEventListener("mouseout", () => {
      hoverElements.forEach((el) => {
        el.classList.remove("hover-color");
      });
    });
  });
}

function setHeight() {
  let nav = document.getElementById("nav");
  let togglediv = document.getElementById("togglediv");
  let querydiv = document.getElementById("querydiv");
  let headers = document.querySelectorAll(".grid-header");
  let dates = document.querySelectorAll(".sticky-date");

  headers.forEach(function (header) {
    header.style.top =
      nav.offsetHeight + togglediv.offsetHeight + querydiv.offsetHeight + "px";
  });
  dates.forEach(function (date) {
    date.style.top =
      nav.offsetHeight +
      togglediv.offsetHeight +
      querydiv.offsetHeight +
      headers[0].offsetHeight +
      20 +
      "px";
  });
}

function stickyDateGridEnd() {
  const dates = document.querySelectorAll(".sticky-date");
  if (dates.length === 0) {
    return;
  }
  if (dates.length === 1) {
    dates[0].parentElement.style.display = "block";
    return;
  }
  let prevDate = "";
  dates.forEach((date) => {
    if (prevDate === "") {
      date.parentElement.style.display = "block";
      prevDate = date;
    } else if (prevDate.textContent === date.textContent) {
      date.parentElement.remove();
      prevDate.parentElement.style.gridRowEnd =
        date.parentElement.style.gridRowEnd;
    } else {
      date.parentElement.style.display = "block";
      prevDate = date;
    }
  });
}
