/* global XMLHttpRequest */
'use strict';

const delay = (() => {
    let timer = 0;
    return (callback, ms) => {
        clearTimeout(timer);
        timer = setTimeout(callback, ms);
    };
})();

// TODO :: Something weird happens when searching on and off and it just stops responding without errors (update; seems to happen when hitting backspace on the last remaining character)

const searchAction = e => {
    let k = e.keyCode;
    // Ignore arrow keys, space, alt, ctrl, cmd and tab key presses
    if ((k >= 37 && k <= 40) || k === 32 || k === 91 || k === 17 || k === 18 || k === 9) return;
    // TODO :: Don't show already selected value(s) in result-list
    // TODO :: Pressing key-down in the search-field should jump into the results-list?
    const t = e.target;
    if (t.dataset.s === '1') return;
    t.dataset.s = '1';
    const rE = t.parentNode.querySelector('.searchResults');
    const l = t.value.trim().length;
    if (l < 2) {
        // Submit the form on enter press on an empty field
        if (l === 0) {
            rE.closest('form').submit();
            return;
        }
        rE.textContent = '';
        rE.style.display = 'none';
        return;
    };

    const noSearchResult = rE.parentNode.querySelector('.noSearchResults');

    // Local search or XHR
    if (null === t.getAttribute('url')) {
        rE.textContent = '';
        const val = t.value.toLowerCase();
        const words = val.replace(/\s{2,}/, ' ').split(' ');
        const options = t.parentNode.querySelector('select').querySelectorAll('option');
        let matches = 0;

        for (let k = 0; k < options.length; k++) {
            let o = options[k];
            let entry = o.innerHTML.toLowerCase();
            if (o.value === '') continue; // Skip `Make a choice`/`Clear selection` field
            let found = 0;
            for (let i = 0; i < words.length; i++) {
                if (entry.indexOf(words[i]) > -1) {
                    found++;
                }
            }
            if (words.length === found) {
                let d = document.createElement('div');
                d.dataset.id = o.value;
                d.appendChild(document.createTextNode(o.innerHTML));
                rE.appendChild(d);
                matches++;
            };
        };
        noSearchResult.style.display = matches > 0 ? 'none' : 'block';
        rE.style.display = matches > 0 ? 'block' : 'none';
        delete t.dataset.s;
    } else {
        const r = new XMLHttpRequest();
        r.open('POST', t.getAttribute('url') + '?s=' + encodeURIComponent(t.value));
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState === 4) {
                rE.textContent = '';
                let rsS = r.responseText.split('\n');
                if (rsS.length === 1 && rsS[0].trim().length === 0) rsS = [];
                for (let k = 0; k < rsS.length; k++) {
                    let entry = rsS[k].split('\t');
                    let d = document.createElement('div');
                    d.dataset.id = entry[0];
                    d.appendChild(document.createTextNode(entry[1]));
                    rE.appendChild(d);
                };
                noSearchResult.style.display = rsS.length > 0 ? 'none' : 'block';
                rE.style.display = rsS.length > 0 ? 'block' : 'none';
            };
            delete t.dataset.s;
        };
    }
};

const delayAction = e => {
    delay(() => {
        searchAction(e);
    }, 350);
};

const selectSearches = document.querySelectorAll('input[search-for]');
for (let i = 0; i < selectSearches.length; i++) {
    selectSearches[i].addEventListener('keydown', e => {
        // Prevent enter press to submit
        if (e.keyCode === 13) e.preventDefault();
    });
    selectSearches[i].addEventListener('keyup', delayAction);
};

const clearResult = e => {
    const b = e.target.parentNode;
    const sb = b.parentNode;
    const sS = sb.parentNode.querySelector('.selectSearch');
    sS.querySelector('option[value="' + b.dataset.id + '"]').removeAttribute('selected');
    sS.selectedIndex = -1;
    sb.removeChild(b);
    sb.parentNode.querySelector('input').style.display = 'block';
};

const searchResultClick = e => {
    const t = e.target;
    const isXHR = t.parentNode.parentNode.querySelector('input').getAttribute('url') !== null;
    const sE = t.parentNode.parentNode.querySelector('select');

    if (isXHR) {
        if (!sE.multiple) sE.textContent = '';
        const o = document.createElement('option');
        o.value = t.dataset.id;
        o.selected = true;
        sE.appendChild(o);
    } else {
        if (!sE.multiple) {
            const sel = sE.querySelector('option[selected]');
            if (null !== sel) sel.selected = false;
        };
        sE.querySelector('option[value="' + t.dataset.id + '"]').selected = true;
    };

    if (!sE.multiple) sE.parentNode.querySelector('input[search-for]').style.display = 'none';

    sE.parentNode.querySelector('.searchResults').style.display = 'none';

    let alreadySelected = false;
    let svE;
    const svEs = sE.parentNode.querySelectorAll('.selectedBadge');
    for (let k = 0; k < svEs.length; k++) {
        svE = svEs[k];
        if (null !== svE.dataset.id && svE.dataset.id === t.dataset.id) {
            alreadySelected = true;
            break;
        };
    };

    if (!alreadySelected) {
        const newSvE = svE.cloneNode(true);
        newSvE.dataset.id = t.dataset.id;
        newSvE.querySelector('.selectedValue').innerHTML = t.innerHTML;
        svE.parentNode.appendChild(newSvE);
        newSvE.style.display = 'inline-block';
        newSvE.querySelector('.clearSelectedValue').addEventListener('click', clearResult);

        sE.parentNode.querySelector('input').value = '';
    };
};

const resultBoxes = document.querySelectorAll('.searchResults');
for (let i = 0; i < resultBoxes.length; i++) resultBoxes[i].addEventListener('click', searchResultClick);

const clearButtons = document.querySelectorAll('.clearSelectedValue');
for (let i = 0; i < clearButtons.length; i++) clearButtons[i].addEventListener('click', clearResult);
