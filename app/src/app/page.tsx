'use client'

import { useState } from 'react'
import { Task, CreateTaskRequest } from '@/lib/api'
import { useTasks } from '@/hooks/useTasks'
import TaskForm from '@/components/TaskForm'
import TaskItem from '@/components/TaskItem'

export default function Home() {
  const [showForm, setShowForm] = useState(false)
  const { tasks, loading, error, createTask, updateTask, deleteTask, refreshTasks, clearError } = useTasks()

  const handleCreateTask = async (taskData: CreateTaskRequest) => {
    try {
      await createTask(taskData)
      setShowForm(false)
    } catch (err) {
      // エラーはuseTasks内で処理される
    }
  }

  const handleUpdateTask = async (id: number, updates: Partial<{ content: string; status: Task['status']; due_date?: string }>) => {
    try {
      await updateTask(id, updates)
    } catch (err) {
      // エラーはuseTasks内で処理される
    }
  }

  const handleDeleteTask = async (id: number) => {
    try {
      await deleteTask(id)
    } catch (err) {
      // エラーはuseTasks内で処理される
    }
  }

  const getTasksByStatus = (status: Task['status']) => {
    return tasks.filter(task => task.status === status)
  }

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-64">
        <div className="text-lg">読み込み中...</div>
      </div>
    )
  }

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-900">タスク管理</h1>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          {showForm ? 'キャンセル' : 'タスクを追加'}
        </button>
      </div>

      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4 flex justify-between items-center">
          <span>{error}</span>
          <button
            onClick={clearError}
            className="text-red-500 hover:text-red-700 text-sm font-bold"
          >
            ×
          </button>
        </div>
      )}

      {showForm && (
        <div className="mb-8">
          <TaskForm
            onSubmit={handleCreateTask}
            onCancel={() => setShowForm(false)}
          />
        </div>
      )}

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <h2 className="text-xl font-semibold mb-4 text-yellow-600">保留中</h2>
          <div className="space-y-4">
            {getTasksByStatus('pending').map(task => (
              <TaskItem
                key={task.id}
                task={task}
                onUpdate={handleUpdateTask}
                onDelete={handleDeleteTask}
              />
            ))}
            {getTasksByStatus('pending').length === 0 && (
              <div className="text-gray-500 text-center py-8">タスクがありません</div>
            )}
          </div>
        </div>

        <div>
          <h2 className="text-xl font-semibold mb-4 text-blue-600">進行中</h2>
          <div className="space-y-4">
            {getTasksByStatus('in_progress').map(task => (
              <TaskItem
                key={task.id}
                task={task}
                onUpdate={handleUpdateTask}
                onDelete={handleDeleteTask}
              />
            ))}
            {getTasksByStatus('in_progress').length === 0 && (
              <div className="text-gray-500 text-center py-8">タスクがありません</div>
            )}
          </div>
        </div>

        <div>
          <h2 className="text-xl font-semibold mb-4 text-green-600">完了</h2>
          <div className="space-y-4">
            {getTasksByStatus('completed').map(task => (
              <TaskItem
                key={task.id}
                task={task}
                onUpdate={handleUpdateTask}
                onDelete={handleDeleteTask}
              />
            ))}
            {getTasksByStatus('completed').length === 0 && (
              <div className="text-gray-500 text-center py-8">タスクがありません</div>
            )}
          </div>
        </div>
      </div>

      <div className="mt-8 text-center">
        <button
          onClick={refreshTasks}
          className="bg-gray-500 text-white px-4 py-2 rounded-md hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500"
          disabled={loading}
        >
          {loading ? '更新中...' : '更新'}
        </button>
      </div>
    </div>
  )
}