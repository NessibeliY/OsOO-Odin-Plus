document.addEventListener("DOMContentLoaded", function() {
    fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1')
        .then(response => response.json())
        .then(data => {
            const cryptoTable = document.getElementById('cryptoTable');
            const tbody = cryptoTable.querySelector('tbody');

            data.forEach((currency, index) => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${currency.id}</td>
                    <td>${currency.symbol}</td>
                    <td>${currency.name}</td>
                `;
                if (currency.symbol === 'usdt') {
                    tr.classList.add('green-bg');
                } else if (index < 5) {
                    tr.classList.add('blue-bg');
                } else {
                    tr.classList.add('white-bg');
                }
                tbody.appendChild(tr);
            });
        })
        .catch(error => console.error('Error fetching data:', error));
});
