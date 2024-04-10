window.onload = (e) => {
  const formTemplate = document.getElementById("form-template");
  const mainForm = document.getElementById("main-form");
  const addFormBtn = document.getElementById("add-form");
  let removeBtns;
  addFormBtn.onclick = () => {
    mainForm.insertBefore(
      formTemplate.content.cloneNode(true),
      mainForm.lastChild
    );
    removeBtns = document.getElementsByClassName("rm-form");
    console.log(removeBtns);
    for (let b of removeBtns) {
      b.onclick = (e) => {
        e.target.parentElement.remove();
      };
    }
  };
};

function addQuery(index, query) {
  const formTemplate = document.getElementById("form-template");
  const mainForm = document.getElementById("main-form");
  const q = formTemplate.content.cloneNode(true);
  const select = q.firstChild.children[0];
  const input = q.firstChild.children[1];
  const btn = q.firstChild.children[2];
  btn.onclick = (e) => {
    e.target.parentElement.remove();
  };
  select.value = index;
  input.value = query;
  mainForm.insertBefore(q, mainForm.lastChild);
  htmx.trigger("#main-form", "onLoadC");
}
