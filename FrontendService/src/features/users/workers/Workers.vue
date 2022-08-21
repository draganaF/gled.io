<template>

  <div class="content">
    <Card title="Workers">

      <SearchWorkerForm @search="onSearchParamsChange($event)" :existingParams="searchParams" />

      <WorkersTable :workers="workers" />
    </Card>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import Card from '../../../components/Card/Card.vue';
import WorkersTable from './WorkersTable.vue';
import SearchWorkerForm from './SearchWorkerForm.vue';

const initialSearchParams = {
  Name: '',
  LastName: '',
  Email: '',
  Role: 1,
  FromNumberOfSoldTickets: 0,
  ToNumberOfSoldTickets: 0
}

export default {
  components: { Card, WorkersTable, SearchWorkerForm },

  data: () => ({
    searchParams: { ...initialSearchParams }
  }),

  computed: {
    ...mapGetters({
      workers: 'users/getUsers',
      result: 'users/getResult'
    })
  },

  watch: {
    result({ label, ok }) {
      if (!ok) return;

      if (['create', 'delete'].includes(label)) {
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