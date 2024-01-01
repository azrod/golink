<template>
  <Generic :item="item">
    <template #indicator>
      <button
        class="button is-info is-small is-rounded is-outlined"
        @click="openModal"
      >
        <span class="icon is-small" @click="openModal">
          <i class="fas fa-info"></i>
        </span>
      </button>
      <div class="modal" :class="{ 'is-active': showModal }">
        <div class="modal-background"></div>
        <div class="modal-card">
          <header class="modal-card-head">
            <p class="modal-card-title">Status Information</p>
            <button
              class="delete"
              aria-label="close"
              @click="closeModal">
              <span class="icon is-small">
                <i class="fas fa-info"></i>
              </span>
            </button>
          </header>
          <section class="modal-card-body">
            <!-- Content ... -->
            <p>{{ status }}</p>
          </section>
          <footer class="modal-card-foot">
            <button class="button" @click="closeModal">Close</button>
          </footer>
        </div>
      </div>
    </template>
  </Generic>
</template>

<script>
import service from "@/mixins/service.js";
import Generic from "./Generic.vue";

export default {
  name: "Info",
  mixins: [service],
  props: {
    item: Object,
  },
  components: {
    Generic,
  },
  data: () => ({
    status: null,
  }),
  created() {
    this.fetchStatus();
  },
  methods: {
    fetchStatus: async function () {
      if (this.item.info.enabled) {
        this.status = "online";
      } else {
        this.status = "offline";
      }

      // const method =
      //   typeof this.item.method === "string"
      //     ? this.item.method.toUpperCase()
      //     : "HEAD";

      // if (!["GET", "HEAD", "OPTION"].includes(method)) {
      //   console.error(`Ping: ${method} is not a supported HTTP method`);
      //   return;
      // }

      // this.fetch("/", { method, cache: "no-cache" }, false)
      //   .then(() => {
      //     this.status = "online";
      //   })
      //   .catch(() => {
      //     this.status = "offline";
      //   });
    },
    closeModal: function () {
      this.showModal = false;
    },
    openModal: function () {
      this.showModal = true;
    },
  },
};
</script>

<style scoped lang="scss">
.status {
  font-size: 0.8rem;
  color: var(--text-title);
  white-space: nowrap;
  margin-left: 0.25rem;

  &.online:before {
    background-color: #94e185;
    border-color: #78d965;
    box-shadow: 0 0 5px 1px #94e185;
  }

  &.offline:before {
    background-color: #c9404d;
    border-color: #c42c3b;
    box-shadow: 0 0 5px 1px #c9404d;
  }

  &:before {
    content: " ";
    display: inline-block;
    width: 7px;
    height: 7px;
    margin-right: 10px;
    border: 1px solid #000;
    border-radius: 7px;
  }
}
</style>
