export function LiveIndicator() {
  return (
    <span className="inline-flex items-center gap-2">
      <span className="w-3 h-3 bg-error rounded-full kinetic-pulse" />
      <span className="px-2 py-0.5 bg-error text-white text-[10px] font-black rounded-sm tracking-wider">
        LIVE
      </span>
    </span>
  );
}
