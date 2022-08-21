<template>
  <div>
    <Table>
      <TableHead
        :columnNames="[
          'Id',
          'Movie',
          'Price',
          'Seat',
          'Reserved',
          'Bought',
          '',
        ]"  
      />
      <TableRow v-for="t in tickets" :key="t.id" :values="getTicketsValues(t)" >
          <DropdownMenu>
            <ModalOpener modalBoxId="buyTicketClick" v-if="!t.is_bought">
              <DropdownItem @click="handleBuyTicketClick(t)"
                >Buy
              </DropdownItem>
            </ModalOpener>

          </DropdownMenu>
        
      </TableRow>
    </Table>

    <Modal modalBoxId="buyTicketClick" title="Buy Ticket">
      <div slot="body">
        <p>
          Are you sure that you want to buy this ticket?
        </p>
      </div>
      <div slot="buttons">
        <OptionModalButtons @yes="handleBuyTicket" />
      </div>
    </Modal>
  </div>
</template>
<script>
import Table from "../../../components/Table/Table.vue";
import TableHead from "../../../components/Table/TableHead.vue";
import TableRow from "../../../components/Table/TableRow.vue";
import DropdownMenu from "../../../components/DropdownMenu/DropdownMenu.vue";
import DropdownItem from "../../../components/DropdownMenu/DropdownItem.vue";
import ModalOpener from "../../../components/Modal/ModalOpener.vue";
import Modal from "../../../components/Modal/Modal.vue";
import OptionModalButtons from "../../../components/Modal/OptionModalButtons.vue";
import { mapActions, mapGetters } from "vuex";
import { getUserIdFromToken } from "../../../utils/token";
import toastr from "toastr";
import moment from 'moment'

export default {
  components: {
    Table,
    TableHead,
    TableRow,
    DropdownMenu,
    DropdownItem,
    ModalOpener,
    Modal,
    OptionModalButtons
  },
  props: ["tickets"],
  
  data: () => ({
    selectedTicket: null
  }),

  computed: {
    ...mapGetters({
      result: "tickets/getResult",
      projections: "projections/getProjections",
    })
  },

  watch: {
    result({ message, ok, label }) {
      if (label === "buy") {
        if (ok) {
          toastr.success(message);
        } else {
          toastr.error(message);
        }
      }
    },
  },
  methods: {
    ...mapActions({
      buyTicket: "tickets/buyTicket",
      fetchProjections: "projections/fetchProjections",
      sellTicket: "users/incrementWorkersSoldTickets",
      incrementUsersTickets: "users/incrementUsersBoughtTickets"
    }),

    formatDateTime(date) {
            return moment(date).format("DD.MM.YYYY HH:mm");
    },
    getTicketsValues(ticket) { 
      return [
        ticket.id,
        this.findProjectionForTicket(ticket.projection_id),
        ticket.price,
        ticket.seat,
        ticket.is_reserved ? "Yes" : "No",
        ticket.is_bought ? "Yes" : "No"
      ]
    },
    findProjectionForTicket(projectionId) {
      let p = this.projections.filter(projection => projection.Id == projectionId);
      return p[0].Movie.Name + "  (" + this.formatDateTime(p[0].Slot) + ")"
    },
    handleBuyTicketClick(t) {
      this.selectedTicket = t;
    },

    handleBuyTicket() {
      this.buyTicket(this.selectedTicket.id)
      this.sellTicket(getUserIdFromToken())
      this.incrementUsersTickets(this.selectedTicket.user_id)
    }
  },
  mounted() {
    this.fetchProjections()
  }
}
</script>
