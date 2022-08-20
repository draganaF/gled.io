import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../features/auth/LoginPage.vue'
import store from '../store/index'
import * as tokenUtils from '../utils/token'

const $ = window.$;

Vue.use(VueRouter)

const routes = [
  {
    path: '/auth',
    name: 'LoginPage',
    component: () => import('@/features/auth/LoginPage.vue'),
    meta: {
      layout: 'AuthLayout'
    }
  },
  {
    path: '/projections',
    name: 'Projections',
    component: () => import('@/features/projection/Projections.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  }
]
    
const router = new VueRouter({
  mode: 'history',
  routes
})

router.beforeEach((to, from, next) => {
  const isUserLoggedIn = tokenUtils.isUserLoggedIn();
  
  
  const { authorizedRoles } = to.meta;
  
  if (authorizedRoles) {
      if (!isUserLoggedIn) {
        store.dispatch('authentication/logOut');
        return next({ path: '/auth' });
      }

      const userRole = tokenUtils.getRoleFromToken();
      if (authorizedRoles.length && !authorizedRoles.includes(userRole)) {
          return next({ path: '/' });
      }
  }

  next();
})

router.afterEach((to, from) => {
  setTimeout(() => {
    $('.selectpicker').selectpicker('refresh');
  }, 100);
});


export default router
