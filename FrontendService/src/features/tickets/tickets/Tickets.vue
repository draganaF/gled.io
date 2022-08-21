<template>
  <div class="content">
    <Card title="Tickets">
      <TicketsTable :tickets="tickets" />
    </Card>
  </div>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import Card from "../../../components/Card/Card.vue";
import TicketsTable from "./TicketsTable.vue";
import { getUserIdFromToken, getRoleFromToken } from "../../../utils/token";
import { Roles } from "../../../constants.js";
export default {
  components: { Card, TicketsTable },
  data: () => ({
    Roles,
    user: {},
  }),
  computed: {
    ...mapGetters({
      tickets: "tickets/getTickets",
      result: "tickets/getResult",
    }),
  },
  watch: {
    result({ message, ok, label }) {
      if (label === "buy") {
        if (ok) {
          this.handleFetching()
        }
      }
    },
  },
  methods: {
    ...mapActions({
      fetchTickets: "tickets/fetchTickets",
      fetchUsersTickets: "tickets/fetchUsersTickets",
    }),

    handleFetching() {
      if (getRoleFromToken() === 0) {
        
        this.fetchUsersTickets(getUserIdFromToken());
      }
      else {
        this.fetchTickets();
      }
    }
  },
  mounted() {
    this.handleFetching();
  },
};
</script>
