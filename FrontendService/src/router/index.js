import Vue from 'vue'
import VueRouter from 'vue-router'

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
    component: () => import('@/features/projection/projections/Projections.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  }, 
  {
    path: '/projections/add-new-projection',
    name: 'CreateProjection',
    component: () => import('@/features/projection/manage-projections/ProjectionForm.vue'),
    meta: {
      layout:'AppLayoutMain'
    }
  },
  {
    path: '/movies',
    name: 'Movies',
    component: () => import('@/features/movies/movies/Movies.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/movies/:id/recensions',
    name: 'Movie Recensions',
    component: () => import('@/features/recensions/recensions-list/Recensions.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/recensions/reports',
    name: 'Recension Reports',
    component: () => import('@/features/recensions/reports/Reports.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/movies/add-new-movie',
    name: 'CreateMovie',
    component: () => import('@/features/movies/manage-movies/MovieForm.vue'),
    meta: {
      layout:'AppLayoutMain'
    }
  },
  {
    path: '/movies/update/:id',
    name: 'UpdateMovie',
    component: () => import('@/features/movies/manage-movies/MovieForm.vue'),
    meta: {
      layout:'AppLayoutMain'
    }
  },
  {
    path: '/users/registration',
    name: 'Registration',
    component: () => import('@/features/users/registration/Registration.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/users/activation',
    name: 'Activation',
    component: () => import('@/features/users/activation/Activation.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/users/workers',
    name: 'Workers',
    component: () => import('@/features/users/workers/Workers.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/users/registred',
    name: 'Registred Users',
    component: () => import('@/features/users/users/Users.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/users/:id',
    name: 'Any User Profile',
    component: () => import('@/features/users/user-profile/UserProfile.vue'),
    meta: {
      ayout: 'AppLayoutMain'
    }
  },
  {
    path: '/users/profile',
    name: 'Current User Profile',
    component: () => import('@/features/users/user-profile/UserProfile.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },
  {
    path: '/tickets',
    name: 'Tickets Table',
    component: () => import('@/features/tickets/tickets/Tickets.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },

  {
    path: '/projections/tickets/:id',
    name: 'Hall Diagram',
    component: () => import('@/features/tickets/tickets-reservation/TicketsReservation.vue'),
    meta: {
      layout: 'AppLayoutMain'
    }
  },

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
    $ && $('.selectpicker').selectpicker('refresh');
  }, 100);
});


export default router
