<template>
<div class="container">
  <Header/>
  <Tasks/>
  <Pagination current-page="1" total-pages="5"/>
</div>
</template>

<script>
import Tasks from "@/components/Tasks";
import Header from "@/components/Header";
import Pagination from "@/components/Pagination";
export default {
name: "Dashboard",
  components: {Pagination, Header, Tasks}
}
</script>

<style scoped>
.container {
  margin: 5px;
}
</style>