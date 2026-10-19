document.addEventListener("keydown", function (e) {
    if (e.key !== "/") return;

    const active = document.activeElement;
    if (active && (active.tagName === "INPUT" || active.tagName === "TEXTAREA" || active.isContentEditable)) return;

    e.preventDefault();
    const search = document.querySelector(".search-box");
    if (search) {
        search.focus();
        search.select();
    }
});

document.addEventListener("click", function (e) {
    if (e.target.id === "add-dialog") {
        e.target.close();
    }
});

document.addEventListener("showNotification", function (e) {
    const container = document.getElementById("notification-container");
    if (!container) return;

    const msg = (typeof e.detail === "object" && e.detail !== null)
        ? (e.detail.value || e.detail.message || JSON.stringify(e.detail))
        : e.detail;

    const card = document.createElement("div");
    card.className = "notification-card";
    card.textContent = msg;

    container.appendChild(card);

    requestAnimationFrame(() => {
        card.classList.add("show");
    });

    setTimeout(() => {
        card.classList.remove("show");
        card.addEventListener("transitionend", () => {
            card.remove();
        });
    }, 3500);
});