'use strict';

let path = window.location.href.split('/')[4];
if (undefined !== path) {
    const a = document.querySelector('.menu a[href$="/' + path.split('?')[0] + '"]');
    // TODO :: Not a problem; but will give an error if it can't resolve an element
    //         if the URL prefix is not in the menu.
    a.className = 'active';
    a.closest('ul').style.display = 'block';
};

document.querySelector('.menu').addEventListener('click', e => {
    const p = e.target;
    if (p.tagName === 'P') {
        if (p.nextSibling) {
            const pd = p.nextSibling.style.display;
            p.nextSibling.style.display = pd === 'none' || pd === '' ? 'block' : 'none';
        }
    }
});
