const express = require('express');
const path = require('path');

const app = express();
const port = process.env.PORT || 3000;

let totalPageViews = 0;

const quotes = [
  'Small progress is still progress.',
  'Build fast, learn faster.',
  'Consistency beats intensity.',
  'Simple solutions scale best.',
  'Make it work, then make it better.'
];

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));
app.use('/public', express.static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
  totalPageViews += 1;
  const hour = new Date().getHours();
  let greeting = 'Hello';

  if (hour < 12) greeting = 'Good morning';
  else if (hour < 18) greeting = 'Good afternoon';
  else greeting = 'Good evening';

  res.render('index', {
    greeting,
    serverTime: new Date().toISOString(),
    totalPageViews
  });
});

app.get('/api/quote', (req, res) => {
  const selectedQuote = quotes[Math.floor(Math.random() * quotes.length)];
  res.json({
    quote: selectedQuote,
    servedAt: new Date().toISOString()
  });
});

app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});
