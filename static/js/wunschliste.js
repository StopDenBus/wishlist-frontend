async function updateUI() {
    const start = (currentPage - 1) * wishesPerPage;
    const end = start + wishesPerPage;
    const wishes = await getWishes();
    const aufgeteilteWuensche = wishes.slice(start, end);

    const divWunschListe = document.getElementById('wunschListe');
    divWunschListe.innerHTML = '';

    if (aufgeteilteWuensche.length === 0) {
        divWunschListe.innerHTML = '<div class="div_keine_wuensche">Keine Wünsche verfügbar.</div>';
    } else {
        const existierendeWuensche = aufgeteilteWuensche.map(w => `
            <div class="wunsch_karte">
              <h3>${w.product}</h3>
              <p><strong>Preis:</strong> ${w.price} €</p>
              <p><strong>Priorität:</strong> ${w.priority}</p>
              <a href="${w.url}" target="_blank">Zum Produkt</a>
              <div class="buttons">
                <button onclick="updateWish(${w.id})">✏️ Bearbeiten</button>
                <button onclick="deleteWish(${w.id})">🗑️ Löschen</button>
              </div>
            </div>
        `);

        const nichtGefuellteWuensche = wishesPerPage - aufgeteilteWuensche.length;
        const leereWuensche = Array(nichtGefuellteWuensche).fill(`
            <div class="wunsch_karte empty_karte"></div>
        `);

        divWunschListe.innerHTML = [...existierendeWuensche, ...leereWuensche].join("");
    }
    anzahlSeiten = Math.ceil(listWuensche.length / wishesPerPage);
    document.getElementById("pageInfo").textContent = `Seite ${currentPage} von ${anzahlSeiten}`;
}

function enableEditing() {
    document.getElementById('editId').value = '';
    document.getElementById('product').value = '';
    document.getElementById('preis').value = '';
    document.getElementById('url').value = '';
    document.getElementById('prioritaet').value = 'Mittel';

    document.getElementById('formPanel').style.display = 'block';
    document.getElementById('formTitle').textContent = 'Neuen Wunsch hinzufügen';
}

async function saveWish() {
    const id = document.getElementById('editId').value;
    const product = document.getElementById('product').value;
    const preis = parseFloat(document.getElementById('preis').value);
    const product_url = document.getElementById('url').value;
    const prioritaet = document.getElementById('prioritaet').value;

    const body = JSON.stringify({ product: product, price: preis, url: product_url, priority: prioritaet })

    if (id) {
        const url = apiURL + "/wish/" + id;
        try {
            const response = await fetch(url, { 
                method: "PUT",
                headers: { "Content-Type": "application/json", },
                body: body,
            });
            if (!response.ok) {
                throw new Error('Response status: ${response.status}');
            }
        } catch (error) {
            console.error(error.message);
        }
    } else {
        const url = apiURL + "/wish";
        try {
            const response = await fetch(url, { 
                method: "POST",
                headers: { "Content-Type": "application/json", },
                body: body,
            });
            if (!response.ok) {
                throw new Error('Response status: ${response.status}');
            }
        } catch (error) {
            console.error(error.message);
        }
    }

    updateUI();
    abortEditing();
}

async function updateWish(id) {
    const url = apiURL + "/wish/by_id/" + id;
    let wish = {};
    try {
        const response = await fetch(url, {
            headers: {
                'Accept': 'application/json',
            },            
        });
        if (!response.ok) {
            throw new Error('Response status: ${response.status}');
        }
        wish = await response.json();
    } catch (error) {
        console.error(error.message);
    }

    if (wish) {
        document.getElementById('editId').value = wish.id;
        document.getElementById('product').value = wish.product;
        document.getElementById('preis').value = wish.price;
        document.getElementById('url').value = wish.url;
        document.getElementById('prioritaet').value = wish.priority;

        document.getElementById('formPanel').style.display = 'block';
        document.getElementById('formTitle').textContent = 'Wunsch bearbeiten';
    }
}

function abortEditing() {
    document.getElementById('editId').value = '';
    document.getElementById('formPanel').style.display = 'none';
    document.getElementById('formTitle').textContent = 'Neuer Wunsch';
}

async function deleteWish(id) {
    const url = apiURL + "/wish/by_id/" + id;
    try {
        const response = await fetch(url, { method: "DELETE"});
        if (!response.ok) {
            throw new Error('Response status: ${response.status}');
        }
    } catch (error) {
        console.error(error.message);
    }
    updateUI();
}

function prepareSort(sort_by) {
    const criteria = sort_by.value;
    if (criteria) {
        sortWishes(criteria);
        sort_by.value = ""; // Zurücksetzen, damit erneut gewählt werden kann
    }
}

function sortWishes(sort_by) {
    if (!sort_by) return;

    // Wenn gleiches Kriterium, Richtung umkehren
    if (sorting.sort_by === sort_by) {
        if (sorting.order_by === OrderBy.ASC) {
            sorting.order_by = OrderBy.DSC;
        } else {
            sorting.order_by = OrderBy.ASC;
        }
    } else {
        sorting.sort_by = sort_by;
        sorting.order_by = OrderBy.ASC;
    }

    currentPage = 1;
    updateUI();
}

function nextPage() {
    if (currentPage * wishesPerPage < getWishes().length) {
        currentPage++;
        updateUI();
    }
}

function previousPage() {
    if (currentPage > 1) {
        currentPage--;
        updateUI();
    }
}

async function getWishes() {
    const url = apiURL + "/wishes?";
    let wishes = [];
    try {
        const response = await fetch(url + new URLSearchParams({
            order_by: sorting.order_by,
            sort_by: sorting.sort_by
        }));
        if (!response.ok) {
            throw new Error('Response status: ${response.status}');
        }
        wishes = await response.json();
    } catch (error) {
        console.error(error.message);
    }

    return wishes;
}