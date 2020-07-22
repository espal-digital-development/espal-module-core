<template>
    <div>
        <button @click="toHome">Home</button>
        <button @click="toLogin">Login</button>
    </div>
</template>

<script>
import { authService, routerService } from '../services';

export default {
    setup() {
        const toHome = () => routerService.goToRoute('Home')
        const toLogin = () => authService.goToLoginPage();
        return { toLogin, toHome };
    }
};
</script>
