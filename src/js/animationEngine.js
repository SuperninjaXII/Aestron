// Select the element with anime-action="scale out"
const element = document.querySelector('[anime-action="scale-out"]');

// Check if the element exists before adding the event listener
if (element) {
  element.addEventListener('click', () => {
    element.classList.add('scale-out-center');
  });
}
