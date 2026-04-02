const express = require('express');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 3000;

const products = [
  { id: 1, name: 'Minimalist Backpack', category: 'Accessories', price: 59.99, rating: 4.6, image: 'https://images.unsplash.com/photo-1491637639811-60e2756cc1c7?w=800' },
  { id: 2, name: 'Wireless Earbuds', category: 'Electronics', price: 89.00, rating: 4.4, image: 'https://images.unsplash.com/photo-1588423771073-b8903fbb85b5?w=800' },
  { id: 3, name: 'Smartwatch Pro', category: 'Electronics', price: 149.50, rating: 4.7, image: 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=800' },
  { id: 4, name: 'Classic Sneakers', category: 'Footwear', price: 74.25, rating: 4.3, image: 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=800' },
  { id: 5, name: 'Cotton Hoodie', category: 'Apparel', price: 42.00, rating: 4.5, image: 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=800' },
  { id: 6, name: 'Desk Lamp', category: 'Home', price: 35.99, rating: 4.2, image: 'https://images.unsplash.com/photo-1524758631624-e2822e304c36?w=800' }
];

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));
app.use(express.static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
  res.render('index', { products });
});

app.get('/api/products', (req, res) => {
  const { category = 'all', sort = 'default', q = '' } = req.query;

  let filtered = products.filter((p) =>
    p.name.toLowerCase().includes(q.toLowerCase()) &&
    (category === 'all' || p.category.toLowerCase() === category.toLowerCase())
  );

  if (sort === 'price_asc') filtered = filtered.sort((a, b) => a.price - b.price);
  if (sort === 'price_desc') filtered = filtered.sort((a, b) => b.price - a.price);
  if (sort === 'rating_desc') filtered = filtered.sort((a, b) => b.rating - a.rating);

  res.json(filtered);
});

app.listen(PORT, () => {
  console.log(`E-commerce app running at http://localhost:${PORT}`);
});
