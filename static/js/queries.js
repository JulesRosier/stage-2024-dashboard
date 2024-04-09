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
