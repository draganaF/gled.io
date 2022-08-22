<template>
  <div>
    <Card :title="projection.Movie.Name ">
    <HallDiagram
      v-if="!!projection"
      :key="keyComponent"
      :projectionId="projectionId"
      :dataDiagram="diagram"
      @chairSelected="handleSelectedChairs"
    />
    <div class="buttons">
      <Button class="button1" @click="reserve" v-if="role=='RegisteredUser'">Reserve</Button>
      <Button class="button2" @click="buy">Buy</Button>
    </div>
    
    </Card>
    
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import HallDiagram from "./HallDiagram.vue";
import Button from "../../../components/Form/Button.vue";
import Card from "../../../components/Card/Card.vue"
import { getRoleFromToken, getUserIdFromToken } from "../../../utils/token";
import toastr from "toastr";
export default {
  components: { HallDiagram, Button, Card },

  data: () => ({
    diagram: {},
    projectionId: 1,
    keyComponent: 0,
    selectedChairs: [],
    role: ""
  }),

  computed: {
    ...mapGetters({
      projection: "projections/getProjection",
      result: "tickets/getResult",
      user: "users/getUser",
    }),
  },

  watch: {
    projection() {
      if (!this.projection) return;
      this.diagram = {
        url:
          this.projection.CinemaHallId === 1
            ? "/diagrams/Sala1.drawio.xml"
            : "/diagrams/Sala2.drawio.xml",
      };
      console.log(this.diagram);
      console.log(this.projection);
      this.keyComponent++;
    },

    result({ ok, message, label }) {
      if (label !== "create") return;
      if (ok) {
        toastr.success(message);
        this.$router.push("/projections");
      } else {
        toastr.error(message);
      }
    },
  },

  methods: {
    ...mapActions({
      fetchProjection: "projections/fetchProjectionById",
      fetchUser: "users/fetchUser",
      createTicket: "tickets/createTicket",
    }),

    handleSelectedChairs(chairs) {
      console.log(chairs);
      this.selectedChairs = [...chairs];
    },
    buy() {
      let userId = this.role == "RegisteredUser" ? this.user.Id : -1

      if (
        this.role == "RegisteredUser" &&
        this.projection.Price * this.selectedChairs.length > this.user.Total
      ) {
        return;
      }
      for (let char of this.selectedChairs) {
        let ticket = {
          is_reserved: false,
          is_bought: true,
          is_deleted: false,
          seat: char,
          user_id: userId,
          projection_id: this.projectionId,
          price: this.projection.Price,
        };
        this.createTicket(ticket);
      }
    },
    reserve() {
      for (let char of this.selectedChairs) {
        let ticket = {
          is_reserved: true,
          is_bought: false,
          is_deleted: false,
          seat: char,
          user_id: this.user.Id,
          projection_id: this.projectionId,
          price: this.projection.Price,
        };
        this.createTicket(ticket);
      }
    },
  },

  mounted() {
    this.projectionId = this.$route.params.id 
    this.fetchProjection(this.projectionId);
    this.role = "Worker"
    if (getRoleFromToken() == "RegisteredUser") {
      this.role = "RegisteredUser"
      this.fetchUser(getUserIdFromToken());
    }
  },
};
</script>

<style scoped>
.buttons {
 display: flex;
  justify-content: center;
  align-items: center;
  height: 100px;
}
.button1 {
  width: 100px;
}
.button2 {
  width: 100px;
}
</style>