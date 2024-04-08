htmx.defineExtension("push-url-w-params", {
  onEvent: function (name, e) {
    if (name === "htmx:configRequest") {
      const path = e.target.getAttribute("data-push-url");
      const params = new URLSearchParams(e.detail.parameters).toString();
      const url = `${window.location.origin}${path}?${params}`;
      window.history.pushState({}, "", url);
    }
  },
});
