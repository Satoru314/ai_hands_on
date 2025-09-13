'use client'

import { Task } from '@/lib/api'

interface TaskItemProps {
  task: Task
  onUpdate: (id: number, updates: Partial<{ content: string; status: Task['status']; due_date?: string }>) => void
  onDelete: (id: number) => void
}

const statusLabels = {
  pending: '保留中',
  in_progress: '進行中',
  completed: '完了',
}

const statusColors = {
  pending: 'bg-yellow-100 text-yellow-800',
  in_progress: 'bg-blue-100 text-blue-800',
  completed: 'bg-green-100 text-green-800',
}

export default function TaskItem({ task, onUpdate, onDelete }: TaskItemProps) {
  const handleStatusChange = (newStatus: Task['status']) => {
    onUpdate(task.id, { status: newStatus })
  }

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleString('ja-JP')
  }

  return (
    <div className="bg-white p-4 rounded-lg shadow-md border-l-4 border-blue-500">
      <div className="flex justify-between items-start mb-2">
        <h3 className="text-lg font-medium text-gray-900">{task.content}</h3>
        <button
          onClick={() => onDelete(task.id)}
          className="text-red-500 hover:text-red-700 text-sm"
        >
          削除
        </button>
      </div>

      <div className="flex items-center space-x-4 mb-2">
        <span className={`px-2 py-1 rounded-full text-xs font-medium ${statusColors[task.status]}`}>
          {statusLabels[task.status]}
        </span>

        <select
          value={task.status}
          onChange={(e) => handleStatusChange(e.target.value as Task['status'])}
          className="text-sm border border-gray-300 rounded px-2 py-1 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="pending">保留中</option>
          <option value="in_progress">進行中</option>
          <option value="completed">完了</option>
        </select>
      </div>

      <div className="text-sm text-gray-500 space-y-1">
        <div>作成: {formatDate(task.created_at)}</div>
        <div>更新: {formatDate(task.updated_at)}</div>
        {task.due_date && (
          <div className="text-orange-600">期限: {formatDate(task.due_date)}</div>
        )}
      </div>
    </div>
  )
}