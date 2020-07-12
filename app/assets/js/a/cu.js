'use strict';

const submitButtons = document.querySelectorAll('input[type=submit]');
for (let i = 0; i < submitButtons.length; i++) submitButtons[i].addEventListener('click', submitAction);

function submitAction(e) {
    e.preventDefault();

    const b = e.target;
    const f = b.closest('form');

    if (undefined !== b.dataset.a) f.querySelector('input[name=Action]').value = b.dataset.a;

    f.submit();
};

// Re-submit blocker
document.querySelector('form').addEventListener('submit', e => {
    const f = e.currentTarget;

    if (f.dataset.s) e.preventDefault();

    f.dataset.s = '1';
});
