<template>

  <div class="content">
    <Card title="Reports">
      <ReportsTable :reports="reports" />
    </Card>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import Card from '../../../components/Card/Card.vue';
import ReportsTable from './ReportsTable.vue';

export default {

  components: { Card, ReportsTable },

  computed: {
    ...mapGetters({
      reports: 'reports/getReports',
      reportResult: 'reports/getResult',
      recensionResult: 'recensions/getResult'
    }),
  },

  watch: {
    reportResult({ ok, label }) {
      if(!ok) return;

      ['silent_delete', 'delete'].includes(label) && this.fetchReports();
    },

    recensionResult({ ok, label }) {
      if(!ok) return;

      label === 'delete' && this.fetchReports();
    },
  },

  methods: {
    ...mapActions({
      fetchReports: 'reports/fetchReports'
    }),
  },

  mounted() {
    this.fetchReports();
  }

}
</script>

<style scoped>
</style>