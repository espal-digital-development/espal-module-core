/* global trans */
'use strict';

const d = document;
const selectAllBoxes = d.querySelectorAll('th input[name=check]');
for (let i = 0; i < selectAllBoxes.length; i++) selectAllBoxes[i].addEventListener('change', toggleAll);

function toggleAll(e) {
    const c = e.target;
    const inputs = c.closest('table').querySelectorAll('td input[name=check]');
    let len = inputs.length;
    for (let i = 0; i < len; i++) inputs[i].checked = c.checked;
};

const toggleButtons = d.querySelectorAll('.listAction');
for (let i = 0; i < toggleButtons.length; i++) toggleButtons[i].addEventListener('click', toggleAction);

function toggleAction(e) {
    e.preventDefault();

    const inputs = e.target.parentElement.parentElement.querySelectorAll('table td input[name=check]:checked');

    const len = inputs.length;
    if (!len) {
        window.alert(trans['nothingSelected']);
        return;
    };

    if (!window.confirm(trans['areYouSure'])) return;

    const ids = [];
    for (let i = 0; i < len; i++) {
        const input = inputs[i];
        if (input.checked && input.dataset.id) ids.push(input.dataset.id);
    };

    window.location.href = e.target.href + (e.target.href.includes("?") ? '&' : '?') + 'ids=' + ids.join(',');
}
