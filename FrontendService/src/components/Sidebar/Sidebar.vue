<template>
  <div class="sidebar" data-color="purple" data-background-color="white">
      <sidebar-logo></sidebar-logo>
      <div class="sidebar-wrapper">
        <ul class="nav" v-if="user.role === Roles.Worker">
          <sidebar-item name="Movies" icon="contacts" to="/movies"></sidebar-item>
          <sidebar-item name="Projections" icon="assignment" to="/projections"></sidebar-item>
          <sidebar-item name="Update user balans" icon="event" to="/work-calendar"></sidebar-item>
        </ul>
        <ul class="nav" v-else-if="user.role === Roles.Administrator">
          <sidebar-item name="Home" icon="dashboard" to="/"></sidebar-item>
          <sidebar-item name="Movies" icon="contacts" to="/movies"></sidebar-item>
        </ul>
        <ul class="nav" v-else>
          <sidebar-item name="Home" icon="dashboard" to="/projections"></sidebar-item>
          <sidebar-item name="Movies" icon="local_pharmacy" to="/movies"></sidebar-item>
        </ul>

      </div>
    </div>
</template>

<script>
import SidebarItem from './SidebarItem.vue'
import SidebarLogo from './SidebarLogo.vue'
import {Roles} from '../../constants.js'
import { getUserIdFromToken, getRoleFromToken } from '../../utils/token'

export default {
  data: () => {
    return {
      Roles,
      user: {}
    }
  },
  mounted() {
    this.user = {
      id: getUserIdFromToken(),
      role: getRoleFromToken()
    }
  },
  components: {
    SidebarLogo,
    SidebarItem
  },
  methods: {
    isUnauthorized() {
      return Object.keys(this.Roles).findIndex(role => this.Roles[role] === this.user.role) === -1;
    }
  }
}
</script>

<style>

</style>