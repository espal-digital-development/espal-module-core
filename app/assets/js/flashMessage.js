'use strict';

const fm = document.querySelector('.flash span');
if (null !== fm) {
    fm.addEventListener('click', e => {
        e.target.parentNode.style.display = 'none';
    });
}
