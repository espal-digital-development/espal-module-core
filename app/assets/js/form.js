'use strict';

// Re-submit blocker
document.querySelector('form').addEventListener('submit', e => {
    const f = e.currentTarget;

    if (f.dataset.s) e.preventDefault();

    f.dataset.s = '1';
});

// Remember me text click-switch
const rm = document.querySelector('input[type=checkbox] + span');
if (rm) {
    rm.addEventListener('click', e => {
        const c = e.currentTarget.previousSibling;
        c.checked = !c.checked;
    });
}
