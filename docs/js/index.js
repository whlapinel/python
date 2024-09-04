"use strict";
(() => {
    console.log("Hello World");
    document.addEventListener("htmx:afterOnLoad", () => {
        console.log("event triggered");
        const pageTitles = ["Home", "About", "Contact", "Blog"];
        const currLocation = window.location.pathname;
        console.log(currLocation);
        for (const title of pageTitles) {
            if (currLocation.includes(title.toLowerCase())) {
                console.log(title);
                const titleElement = document.querySelector(`#nav-${title.toLowerCase()}`);
                console.log(titleElement);
                titleElement === null || titleElement === void 0 ? void 0 : titleElement.classList.toggle("bg-green-700");
            }
        }
    });
})();
