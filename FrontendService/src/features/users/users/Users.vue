<template>

  <div class="content">
    <Card title="Users">

      <SearchUserForm @search="onSearchParamsChange($event)" :existingParams="searchParams" />

      <UsersTable :users="users" />
    </Card>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import Card from '../../../components/Card/Card.vue';
import UsersTable from './UsersTable.vue';
import SearchUserForm from './SearchUserForm.vue';

const initialSearchParams = {
  Name: '',
  LastName: '',
  Email: '',
  Role: 0,
  FromNumberOfBoughtTickets: 0,
  ToNumberOfBoughtTickets: 0,
  FromNumberOfResevedTickets: 0,
  ToNumberOfResevedTickets: 0
}

export default {
  components: { Card, UsersTable, SearchUserForm },

  data: () => ({
    searchParams: { ...initialSearchParams }
  }),

  computed: {
    ...mapGetters({
      users: 'users/getUsers',
      result: 'users/getResult'
    })
  },

  watch: {
    result({ label, ok }) {
      if (!ok) return;

      if (['block'].includes(label)) {
        this.handleSearch();
      }
    }
  },

  methods: {
    ...mapActions({
      searchUsers: 'users/searchUsers'
    }),

    onSearchParamsChange(params) {
      this.searchParams = { ...this.searchParams, ...params };
      this.handleSearch();
    },

    handleSearch() {
      this.searchUsers(this.searchParams);
    }
  },

  mounted() {
    this.handleSearch();
  }

}
</script>

<style>
</style>