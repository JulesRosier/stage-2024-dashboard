window.onload = (e) => {
  const formTemplate = document.getElementById("form-template");
  const queries = document.getElementById("queries");
  const addFormBtn = document.getElementById("add-form");
  addFormBtn.onclick = () => {
    queries.appendChild(formTemplate.content.cloneNode(true));
    addL();
    setHeight();
  };
  Sortable.create(queries, { animation: 150, onUpdate: reload });
  addL();
  document.cookie = "timezone=" + Intl.DateTimeFormat().resolvedOptions().timeZone;
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
    header.style.top = nav.offsetHeight + togglediv.offsetHeight + querydiv.offsetHeight + "px";
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

function somethingSticky() {
  let dates = document.querySelectorAll(".sticky-date");
  console.log(length(dates));
  if (dates.length <= 1) {
    return;
  }
  let prevDate = "";
  dates.forEach(function (date) {
    //check if previous date is same as current date, set grid row end to new ending
    //nu alle dates, maar moet enkel de nieuwe dates hebben
    //else make new date
    if (prevDate === "") {
      prevDate = date;
    } else if (prevDate === date.innerText) {
      // vorige sticky date is hetzelfde als de huidige sticky date
      console.log(prevDate.parentElement);
      prevDate.parentElement.style.gridRowEnd = date.parentElement.style.gridRowEnd;
      date.parentElement.remove();
      return;
    } else {
      // vorige sticky date is niet hetzelfde als de huidige sticky date
      return;
    }
  });
}
