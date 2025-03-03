export function NotFound() {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-center">
        <h1 className="font-bold text-8xl text-error mb-4">404</h1>
        <div className="text-5xl font-bold mb-8">Page Not Found</div>
        <p className="text-gray-500 text-xl mb-8">The page you're looking for doesn't exist or has been moved.</p>
        <a href="/" className="btn btn-primary">Back to Home</a>
      </div>
    </div>
  );
}
