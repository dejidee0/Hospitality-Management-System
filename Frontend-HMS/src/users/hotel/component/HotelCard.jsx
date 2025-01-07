import img from "../../../assets/signup-side-image.svg";


export function HotelCard({ name, location, image, rating, reviews, price }) {
  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden">
      <div className="relative h-48">
        <img src={img} alt={name} className="object-none" />
      </div>
      <div className="p-4">
        <h3 className="font-semibold text-lg">{name}</h3>
        <p className="text-sm text-gray-600">{location}</p>
        <div className="flex items-center gap-2 my-2">
          <span className="text-sm font-medium">{rating}</span>
          <span className="text-sm text-gray-600">({reviews} Reviews)</span>
        </div>
        <div className="flex items-center justify-between">
          <div>
            <p className="text-sm text-gray-600">Per Night</p>
            <p className="text-lg font-semibold">₦{price.toLocaleString()}</p>
          </div>
        </div>
      </div>
    </div>
  );
}
