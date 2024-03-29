class JSONViewer extends HTMLElement {
  constructor() {
    super();
  }
  connectedCallback() {
    if (!this.hasAttribute("data-json")) {
      this.setAttribute("data-json", this.innerHTML);
    }

    const pjson = JSON.stringify(
      JSON.parse(atob(this.getAttribute("data-json"))),
      null,
      2
    );

    const coloredContent = pjson

      .replace(
        /"(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?/g,
        (match) => {
          if (/"(\w+)"\s*:/g.test(match)) {
            return '<span class="json-key">' + match + "</span>";
          } else if (/^"/.test(match)) {
            return '<span class="json-string">' + match + "</span>";
          } else if (/true|false/.test(match)) {
            return '<span class="json-boolean">' + match + "</span>";
          } else if (/null/.test(match)) {
            return '<span class="json-null">' + match + "</span>";
          } else {
            return '<span class="json-number">' + match + "</span>";
          }
        }
      )
      .replace(/[{}\[\],]/g, (match) => {
        return '<span class="json-brace">' + match + "</span>";
      });
    this.innerHTML = "<pre>" + coloredContent + "</pre>";
  }
}

customElements.define("json-viewer", JSONViewer);
