<template>
  <div class="app-container">

    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" stripe fit highlight-current-row>
      <el-table-column label="ID" width="95" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="标题" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="描述" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <!-- <el-table-column label="内容" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.content }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图片" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.cover_image_url }}
        </template>
      </el-table-column> -->
      <el-table-column label="状态" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusTypeFilter">{{ scope.row.state | statusDisplayFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="tag" label="文章标签" width="200" align="center">
        <template slot-scope="scope">
          <el-tag v-for="tag in scope.row.tags" :key="tag.id">{{ tag.name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
      <el-table-column label="更新时间" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.update_time }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="Actions" width="120">
        <template slot-scope="scope">
          <router-link :to="'/article/edit/' + scope.row.id">
            <el-button type="primary" size="small" icon="el-icon-edit">
              编辑
            </el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { listArticle } from '@/api/article'

export default {
  filters: {
    statusTypeFilter (status) {
      const statusMap = {
        1: 'success',
        2: 'danger',
        3: 'gray'
      }
      return statusMap[status]
    },
    statusDisplayFilter (status) {
      const statusMap = {
        1: '发布',
        2: '草稿',
        3: '删除'
      }
      return statusMap[status]
    }
  },
  data () {
    return {
      list: null,
      listLoading: true
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.listLoading = true
      listArticle({ state: 1 }).then(response => {
        this.list = response.data
        this.listLoading = false
      })
    }
  }
}
</script>
