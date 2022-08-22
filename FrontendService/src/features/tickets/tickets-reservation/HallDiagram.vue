<template>
  <div class="diagram" id="hallDiagram">
    <div
      class="drawio-diagram"
      :data-diagram-url="dataDiagram.url"
    >
    </div>
  </div>

</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import { initElements } from './diagramsHandler';

export default {

  props: ['dataDiagram', 'projectionId'],

  data: () => ({
    chain: null,

    selectedChairs: [],
    chairStates: {},
    elements: {}
  }),

  computed: {
    ...mapGetters({
      tickets: 'tickets/getTickets'
    })
  },

  watch: {
    dataDiagram() {
      if (!this.dataDiagram) return;
    },
  },

  methods: {
    ...mapActions({
      fetchTicketsForProjection: 'tickets/fetchTicketsByProjectionId'
    }),

    async init() {

      const { elements } = await initElements({
        id: 'hallDiagram',
        idPrefix: '$element_',
        onClick: (element) => {
          const { id: chairIdWithColor } = element;
          const chairId = chairIdWithColor.split('_')[0];

          if (this.chairStates[chairId] === 'Taken') {
            return;
          }

          if (this.chairStates[chairId] === 'NotTaken') {
            this.chairStates = { ...this.chairStates, [chairId]: 'Selected' };
            this.selectedChairs = [...this.selectedChairs, chairId];
            this.elements[`${chairId}_white`].hide();
            this.elements[`${chairId}_green`].show();
          } else {
            this.chairStates = { ...this.chairStates, [chairId]: 'NotTaken' };
            this.selectedChairs = this.selectedChairs.filter(ch => ch !== chairId);
            this.elements[`${chairId}_white`].show();
            this.elements[`${chairId}_green`].hide();
          }

          this.$emit('chairSelected', this.selectedChairs )
        },
      });

      const initialChairStatuses = {};
      for(const chairIdWithColor of Object.keys(elements)) {
        if (chairIdWithColor.includes("_white")) {
          const chairId = chairIdWithColor.split('_')[0];
          initialChairStatuses[chairId] = 'NotTaken';
        } else {
          elements[chairIdWithColor].hide();
        }
      }

      this.tickets.forEach(t => {
      initialChairStatuses[t.seat] = 'Taken';
      elements[`${t.seat}_red`].show();
      elements[`${t.seat}_white`].hide();

      })


      this.chairStates = {...initialChairStatuses};
      this.elements = elements;
      
      return { elements };
    }

  },

  mounted() {
    this.fetchTicketsForProjection(this.projectionId);
    setTimeout(async () => {
      if(!this.dataDiagram || !this.dataDiagram.url) return;
      await this.init();
    }, 1000);

  }
}
</script>

<style>

.diagram {
  width: 100%;
  padding: 0 10vw;
  box-sizing: border-box
}

</style>