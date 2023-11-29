<template>
  <div>
    <vxe-grid ref="xGrid" v-bind="gridOptions" v-on="gridEvent">
      <template #toolbar_buttons>
        <vxe-button status="primary" content="新建组织" @click="AddTask" />
      </template>
      <template #operate="{ row }">
        <vxe-button status="danger" content="删除" @click="removeRowEvent(row)" />
      </template>
    </vxe-grid>
    <el-dialog v-model="dialogFormVisible" title="新建组织">
      <el-form :model="form" label-width="120px">
        <el-form-item label="组织名称">
          <el-input v-model="form.organizationName" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="form.organizationDescription" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmitCreatOrganize">Create</el-button>
          <el-button>Cancel</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog
      v-model="dialogSearchRecord"
      width="30%"
      title="Inner Dialog"
      append-to-body
    >
      <el-table
        ref="singleTableRef"
        :data="tableData"
        highlight-current-row
        style="width: 100%"
        @row-click="handleRowClick"
      >
        <el-table-column type="index" width="50" />
        <el-table-column property="ID" label="ID" width="120" />
        <el-table-column property="name" label="搜索规则" width="120" />
        <el-table-column property="CreatedAt" label="创建时间" width="120" />
        <el-table-column property="UpdatedAt" label="更新时间" width="120" />
        <el-table-column property="description" label="描述" width="120" />
        <el-table-column property="count" label="资产数量" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
// import { getOrganizationList, deleteOrganizationByIds } from '@/api/organization'
import router from '@/router/index'
import { onBeforeUnmount, onMounted, reactive, ref } from 'vue'
import { VXETable, VxeGridInstance, VxeGridListeners, VxeGridProps } from 'vxe-table'
import { ElMessage } from 'element-plus'

interface RowVO {
  [key: string]: any
}

const xGrid = ref<VxeGridInstance<RowVO>>()
const gridOptions = reactive<VxeGridProps<RowVO>>({
  border: true,
  showHeaderOverflow: true,
  showOverflow: true,
  keepSource: true,
  id: 'full_edit_1',
  height: 600,
  rowConfig: {
    keyField: 'id',
    isHover: true
  },
  columnConfig: {
    resizable: true
  },
  editConfig: {
    trigger: 'manual',
    mode: 'row',
    showStatus: true
  },
  customConfig: {
    storage: true,
    checkMethod({ column }) {
      if (['nickname', 'role'].includes(column.field)) {
        return false
      }
      return true
    }
  },
  printConfig: {
    columns: [
      { field: 'name' },
      { field: 'email' },
      { field: 'nickname' },
      { field: 'age' },
      { field: 'amount' }
    ]
  },
  sortConfig: {
    trigger: 'cell',
    remote: true
  },
  filterConfig: {
    remote: true
  },
  pagerConfig: {
    pageSize: 10,
    pageSizes: [5, 10, 15, 20, 50, 100, 200, 500, 1000]
  },
  formConfig: {
    titleWidth: 100,
    titleAlign: 'right',
    items: [
      { field: 'name', title: '搜索规则', span: 8, titlePrefix: { message: 'app.body.valid.rName', icon: 'vxe-icon-question-circle-fill' }, itemRender: { name: '$input', props: { placeholder: '请输入名称' }}},
      { field: 'email', title: '描述', span: 8, titlePrefix: { useHTML: true, message: '点击链接：<a class="link" href="https://vxetable.cn" target="_blank">vxe-table官网</a>', icon: 'vxe-icon-question-circle-fill' }, itemRender: { name: '$input', props: { placeholder: '请输入邮件' }}},
      { field: 'begintime', title: '开始时间', span: 8, itemRender: { name: '$input', props: { placeholder: '请输入昵称' }}},
      { field: 'endtime', title: '最后时间', span: 8, itemRender: { name: '$input', props: { placeholder: '请输入昵称' }}},
      { span: 24, align: 'center', collapseNode: true, itemRender: { name: '$buttons', children: [{ props: { type: 'submit', content: '搜索', status: 'primary' }}, { props: { type: 'reset', content: '清空' }}] }}
    ]
  },
  toolbarConfig: {
    custom: true,
    slots: {
      buttons: 'toolbar_buttons',
      // tools: 'toolbar_tools'
    },
    import: true,
    export: true,
    zoom: true,
  },
  proxyConfig: {
    seq: true, // 启用动态序号代理，每一页的序号会根据当前页数变化
    sort: true, // 启用排序代理，当点击排序时会自动触发 query 行为
    filter: true, // 启用筛选代理，当点击筛选时会自动触发 query 行为
    form: true, // 启用表单代理，当点击表单提交按钮时会自动触发 reload 行为
    props: {
      // 对应响应结果 Promise<{ result: [], page: { total: 100 } }>
      result: 'data.list', // 配置响应结果列表字段
      total: 'page.total' // 配置响应结果总页数字段
    },
    // 只接收Promise，具体实现自由发挥
    ajax: {
      // 当点击工具栏查询按钮或者手动提交指令 query或reload 时会被触发
      query: ({ page, sorts, filters, form }) => {
        const queryParams: any = Object.assign({}, form) // 创建查询参数
        // 处理排序条件
        const firstSort = sorts[0]
        if (firstSort) {
          queryParams.sort = firstSort.field
          queryParams.order = firstSort.order
        }
        // 处理筛选条件
        filters.forEach(({ field, values }) => {
          queryParams[field] = values.join(',')
        })
        return null
        // console.log(getSearchRecordList({ page: 1, pageSize: 10 }).then(response => response.data.list))
        // return getOrganizationList()
      },
    }
  },
  columns: [
    { type: 'checkbox', title: '状态', width: 20 },
    { field: 'ID', title: '序号', sortable: true, width: 20 },
    { field: 'name', title: '组织名称', sortable: true, titlePrefix: { message: '名称必须填写！' }},
    { field: 'description', title: '组织描述', width: 160 },
    { field: 'CreatedAt', title: '创建时间', width: 160 },
    { field: 'UpdatedAt', title: '更新时间' },
    { title: '操作', width: 300, slots: { default: 'operate' }}
  ],
})

const hasActiveEditRow = (row: RowVO) => {
  const $grid = xGrid.value
  if ($grid) {
    return $grid.isEditByRow(row)
  }
  return false
}

const editRowEvent = (row: RowVO) => {
  const $grid = xGrid.value
  if ($grid) {
    $grid.setEditRow(row)
  }
}

const clearRowEvent = () => {
  const $grid = xGrid.value
  if ($grid) {
    $grid.clearEdit()
  }
}

const saveRowEvent = async(row: RowVO) => {
  const $grid = xGrid.value
  if ($grid) {
    await $grid.clearEdit()
    gridOptions.loading = true
    // 模拟异步保存
    setTimeout(() => {
      gridOptions.loading = false
      VXETable.modal.message({ content: `${JSON.stringify(row)}`, status: 'success' })
    }, 300)
  }
}

const removeRowEvent = async(row: RowVO) => {
  const type = await VXETable.modal.confirm('您确定要删除该数据?')
  const $grid = xGrid.value
  if ($grid) {
    if (type === 'confirm') {
      const ids = []
      ids.push(row.ID)
      const result = await deleteOrganizationByIds({ ids: ids })
      if (result.code === 0) {
        ElMessage.info('删除成功')
        await xGrid.value.commitProxy('query') // 手动提交query查询
      } else {
        ElMessage.error('删除成功')
      }
    }
  }
}

const gridEvent: VxeGridListeners<RowVO> = {
  cellDblclick({ row }) {
    router.push({
      name: 'organizeDetails',
      params: {
        id: row.ID,
      },
    })
  }
}

// 添加任务
// import {
//   createOrganization
// } from '@/api/organization'
const dialogFormVisible = ref(false)
const dialogSearchRecord = ref(false)
const onSubmitCreatOrganize = async() => { // 提交任务
  const result = await createOrganization({
    'name': form.organizationName,
    'description': form.organizationDescription
  })
  if (result.code === 0) {
    ElMessage.info('添加成功')
    await xGrid.value.commitProxy('query') // 手动提交query查询
  } else {
    ElMessage.error('添加失败')
  }
}
// 表单绑定的值
const form = reactive({
  organizationName: '', // 组织名字
  organizationDescription: '', // 组织类型
})
const AddTask = () => { // 打开任务添加弹窗
  dialogFormVisible.value = true // 打开弹窗
}
// 搜索记录操作
// import {
//   getSearchRecordList
// } from '@/api/searchRecord'
let tableData = [ // 历史查询记录数据
]
const SelectSearchRecord = async() => { // 打开选择搜索记录的弹窗
  const result = await getSearchRecordList()
  if (result.code === 0) {
    tableData = result.data.list
  }
  dialogSearchRecord.value = true
}

const currentRow = ref()
const selectRecord = ref({
  id: '',
  name: ''
})
const handleRowClick = (row, column, event) => { // 将选中的值展示到输入框，然后退出弹框
  selectRecord.value.id = row.ID
  selectRecord.value.name = row.name
  dialogSearchRecord.value = false
}

// 初始化函数
onMounted(async() => {

  // const list = await findSearchRecord({page: 1, pageSize: 10})
})

// 在组件卸载之前移除侦听
onBeforeUnmount(() => {
  console.log('卸载searchRecord')
  // emitter.off('sayHi')
})
</script>
