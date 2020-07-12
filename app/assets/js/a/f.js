'use strict';

function urlParam(u, n, v) {
    const pattern = new RegExp('\\b(' + n + '=).*?(&|$)');

    // TODO :: This works, but in does weird modifications like `Userr=100&` going from page 2 to page 1
    // if (v === '') {
    //     u = u.replace(pattern, '');
    //     if (u.indexOf('?') > 0) return u.replace('?', '');
    //     return u;
    // };

    if (u.search(pattern) >= 0) return u.replace(pattern, '$1' + v + '$2');
    u = u.replace(/\?$/, '');
    return u + (u.indexOf('?') > 0 ? '&' : '?') + n + '=' + v;
};

const pSelect = document.querySelector('.pagination select');
if (null !== pSelect) {
    pSelect.addEventListener('change', e => {
        window.location.href = urlParam(urlParam(window.location.href, 'p', ''), 'r', e.target.value);
    });
};

const pages = document.querySelectorAll('.pagination p');
for (let i = 0; i < pages.length; i++) pages[i].addEventListener('click', clickPage);
function clickPage(e) {
    const p = e.target.innerHTML;
    window.location.href = urlParam(window.location.href, 'p', p === '1' ? '' : p);
    return false;
};

const search = document.getElementById('filterSearch');
if (null !== search) {
    search.addEventListener('keyup', e => {
        if (e.keyCode !== 13) return;

        const v = e.target.value.trim();
        let url = window.location.href;

        // Don't trigger if there wasn't anything searched for already
        if (v.length === 0 && url.indexOf('s=') === -1) return;

        url = urlParam(urlParam(url, 'p', ''), 's', v);

        window.location.href = v.length === 0 ? url.replace('&s=', '') : url;
    });
}
