<template>
  <div class="container">
    <div class="title">
      <h1>Current Tasks</h1>
    </div>
    <div class="page">
      <table class="table">
        <tr class="tableHeader">
          <th class="tableHead">Name</th>
          <th class="tableHead">Project</th>
          <th class="tableHead tableHeadCentered">Status</th>
          <th class="tableHead tableHeadCentered">Commits</th>
        </tr>
        <tr class="tableRow" v-for="item in tasks" :key="item.id">
          <td class="tableCell">
            {{ item.name }}
          </td>
          <td class="tableCell">
            {{ item.project.name }}
          </td>
          <td class="tableCell tableCellCentered">
            {{ item.status }}
          </td>
          <td class="tableCell tableCellCentered">
            {{ item.commits }}
          </td>
          <td class="tableCell tableCellAction">
            <feather class="tableActionButton" type="arrow-right"/>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "Tasks",
  data: function() {
    return {
      tasks: [
        { id: 0, name: 'task test', status: 'Ready To Start', commits: 5, project: { name: 'Some Project Name'}},
        { id: 1, name: 'task test', status: 'In Progress', commits: 5, project: { name: 'Some Project Name'}},
        { id: 2, name: 'task test', status: 'Completed', commits: 5, project: { name: 'Some Project Name'}},
      ]
    }
  },
}
</script>

<style scoped>
.title {
  padding-left: 10%;
  text-align: left;
}

.page {
  width: 100%;
}

.table {
  margin: 0 auto;
  width: 80%;
  position: relative;
  border-spacing: 0 10px;
  border-collapse: separate;
}

.tableHead {
  color: dimgrey;
  font-size: 10pt;
  font-weight: normal;
  font-family: "Roboto", sans-serif;
  box-sizing: border-box;
  text-align: left;
  padding-left: 35px;
}

.tableHeadCentered {
  text-align: center;
}

.tableRow {
  width: 100%;
  margin-bottom: 6px;
  border: 1px solid #555;
}

td {
  border: 1px solid #555;
  border-style: solid none;
}

td:first-child {
  border-left-style: solid;
  border-top-left-radius: 6px;
  border-bottom-left-radius: 6px;
}

td:last-child {
  border-right-style: solid;
  border-bottom-right-radius: 6px;
  border-top-right-radius: 6px;
}

.tableCell {
  color: white;
  font-size: 11pt;
  font-weight: normal;
  font-family: "Roboto", sans-serif;
  box-sizing: border-box;
  height: 50px;
  background: dimgrey;
  text-align: left;
  padding-left: 35px;
}

.tableCellCentered {
  text-align: center;
  /*padding-left: 50px;*/
}

.tableCellAction {
  color: white;
  text-align: right;
  padding-right: 15px;
}

.tableCellAction:hover {
  color: darkgoldenrod;
}

.tableActionButton {
  height: 20px;
  width: 20px;
  color: white;
  cursor: pointer;
}
</style>