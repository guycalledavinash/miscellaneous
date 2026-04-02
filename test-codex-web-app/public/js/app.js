const grid = document.getElementById('product-grid');
const cartCountElement = document.getElementById('cart-count');
const searchInput = document.getElementById('search');
const categorySelect = document.getElementById('category');
const sortSelect = document.getElementById('sort');

let cart = JSON.parse(localStorage.getItem('cart') || '[]');
let currentProducts = window.__INITIAL_PRODUCTS__ || [];

function updateCartUI() {
  cartCountElement.textContent = cart.length;
  localStorage.setItem('cart', JSON.stringify(cart));
}

function renderProducts(products) {
  grid.innerHTML = products.map((product) => `
    <article class="product-card" data-id="${product.id}">
      <img src="${product.image}" alt="${product.name}" loading="lazy" />
      <h3>${product.name}</h3>
      <p>${product.category}</p>
      <p>⭐ ${product.rating}</p>
      <p class="price">$${product.price.toFixed(2)}</p>
      <button class="add-to-cart" data-id="${product.id}">Add to Cart</button>
    </article>
  `).join('');
}

async function fetchProducts() {
  const params = new URLSearchParams({
    q: searchInput.value,
    category: categorySelect.value,
    sort: sortSelect.value,
  });

  const response = await fetch(`/api/products?${params.toString()}`);
  currentProducts = await response.json();
  renderProducts(currentProducts);
}

grid.addEventListener('click', (e) => {
  if (!e.target.classList.contains('add-to-cart')) return;

  const productId = Number(e.target.dataset.id);
  const product = currentProducts.find((p) => p.id === productId) ||
    (window.__INITIAL_PRODUCTS__ || []).find((p) => p.id === productId);

  if (product) {
    cart.push(product);
    updateCartUI();
  }
});

[searchInput, categorySelect, sortSelect].forEach((el) => {
  el.addEventListener('input', fetchProducts);
  el.addEventListener('change', fetchProducts);
});

updateCartUI();
