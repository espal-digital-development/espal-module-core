import {h} from 'vue';

export default (width, height, onclick) =>
    h('img', {
        src: require('../../images/logo.png').default.replace('js/spa/images', 'i'),
        width,
        height,
        onclick
    });
