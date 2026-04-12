function SkeletonRow() {
    return (
      <div className="flex items-center gap-3 p-4 bg-white rounded-lg border border-gray-100">
        <div className="w-5 h-5 rounded bg-gray-200 animate-pulse flex-shrink-0" />
        <div className="flex-1 h-4 rounded bg-gray-200 animate-pulse" />
        <div className="w-4 h-4 rounded bg-gray-200 animate-pulse" />
      </div>
    )
  }
  
  export default function TodoSkeleton() {
    return (
      <div className="flex flex-col gap-2">
        {[...Array(3)].map((_, i) => <SkeletonRow key={i} />)}
      </div>
    )
  }