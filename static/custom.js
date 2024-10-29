document.addEventListener("DOMContentLoaded", function () {
  // Find all modified links
  const modifiedLinks = document.querySelectorAll('a[href*="m-wikipedia.org"]');

  // Apply effects to each link
  modifiedLinks.forEach((link) => {
    // Add hover effect
    link.addEventListener("mouseenter", function (e) {
      this.classList.add("magic-effect");
    });

    // Remove effect after animation ends
    link.addEventListener("animationend", function (e) {
      this.classList.remove("magic-effect");
    });

    // Add click effect
    link.addEventListener("click", function (e) {
      e.preventDefault();
      this.style.transform = "scale(0.95)";
      setTimeout(() => {
        this.style.transform = "scale(1)";
        window.location.href = this.href;
      }, 200);
    });
  });

  // Log modification info
  console.log("***Modified links:", modifiedLinks.length);
});
