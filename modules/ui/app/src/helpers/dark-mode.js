export function darkMode(value) {
  if (value === 'dark') {
    document.body.setAttribute('data-bs-theme', value);
  } else if (value === 'auto') {
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
      document.body.setAttribute('data-bs-theme', value);
    } else {
      document.body.removeAttribute('data-bs-theme');
    }
  } else {
    document.body.removeAttribute('data-bs-theme');
  }
}