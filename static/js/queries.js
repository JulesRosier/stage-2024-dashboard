window.onload = (e) => {
  const formTemplate = document.getElementById("form-template");
  const queries = document.getElementById("queries");
  const addFormBtn = document.getElementById("add-form");
  addFormBtn.onclick = () => {
    queries.appendChild(formTemplate.content.cloneNode(true));
    addL();
  };
  Sortable.create(queries, { animation: 150, onUpdate: reload });
  addL();
};

function addL() {
  let removeBtns;
  removeBtns = document.getElementsByClassName("rm-form");
  for (let b of removeBtns) {
    b.onclick = (e) => {
      e.target.parentElement.parentElement.remove();
      reload();
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
