import React, { useState, useEffect } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

const SearchBar = () => {
  const [dropdown, setDropdown] = useState({
    whereTo: false,
    guests: false
  });
  const [rooms, setRooms] = useState([
    { rooms: 1, adults: 1, children: 0 } // Initial room
  ]);
  const [values, setValues] = useState({
    whereTo: "",
    checkIn: null,
    checkOut: null,
    guests: { adults: 1, children: 0 }
  });

  const [suggestions, setSuggestions] = useState([]);

  useEffect(() => {
    const totalGuests = rooms.reduce(
      (totals, room) => {
        totals.adults += room.adults;
        totals.children += room.children;
        return totals;
      },
      { adults: 0, children: 0 }
    );
  
    // Update `guests` with the total guests and the rooms count
    setValues(prev => ({ 
      ...prev, 
      guests: { 
        ...totalGuests, 
        rooms: rooms.length 
      } 
    }));
  }, [rooms]);

  const showResult = () => {
    console.log(values);
  };

  const updateRoomCount = (index, type, operation) => {
    setRooms(prev =>
      prev.map((room, i) =>
        i === index
          ? {
              ...room,
              [type]:
                operation === "increment"
                  ? room[type] + 1
                  : Math.max(room[type] - 1, 0)
            }
          : room
      )
    );
  };

  const addRoom = () => {
    setRooms(prev => [
      ...prev,
      { rooms: prev.length + 1, adults: 1, children: 0 }
    ]);
  };

  const removeRoom = index => {
    setRooms(prev =>
      prev
        .filter((_, i) => i !== index)
        .map((room, i) => ({ ...room, rooms: i + 1 }))
    );
  };

  const handleInputChange = async e => {
    const input = e.target.value;
    setValues(prev => ({ ...prev, whereTo: input }));

    if (input.trim().length > 0) {
      try {
        const response = await fetch(`/api/locations?query=${input}`);
        const data = await response.json();
        setSuggestions(data);
        setDropdown(prev => ({ ...prev, whereTo: true }));
      } catch (error) {
        console.error("Error fetching suggestions:", error);
        setSuggestions([]);
      }
    } else {
      setDropdown(prev => ({ ...prev, whereTo: false }));
      setSuggestions([]);
    }
  };

  const selectOption = (field, option) => {
    setValues(prev => ({ ...prev, [field]: option }));
    setDropdown(prev => ({ ...prev, whereTo: false }));
    setSuggestions([]);
  };

  const totalGuests = rooms.reduce(
    (totals, room) => {
      totals.adults += room.adults;
      totals.children += room.children;
      return totals;
    },
    { adults: 0, children: 0 }
  );

  const totalRooms = rooms.length;

  return (
    <div className="flex gap-4 items-center p-4 bg-white shadow-lg rounded-lg">
      {/* Where To */}
      <div className="relative w-1/4">
        <label className="block text-xs text-gray-500 mb-1">Where To</label>
        <input
          type="text"
          className="border rounded-md px-4 py-2 w-full text-gray-700"
          placeholder="Enter location"
          value={values.whereTo}
          onChange={handleInputChange}
          onFocus={() => setDropdown(prev => ({ ...prev, whereTo: true }))}
          onBlur={() =>
            setTimeout(() => setDropdown(prev => ({ ...prev, whereTo: false })), 200)
          }
        />
        {dropdown.whereTo && suggestions.length > 0 && (
          <div className="absolute bg-white border rounded-md mt-1 shadow-lg w-full z-10">
            {suggestions.map((option, index) => (
              <div
                key={index}
                className="px-4 py-2 cursor-pointer hover:bg-gray-100"
                onClick={() => selectOption("whereTo", option)}
              >
                {option}
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Check In */}
      <div>
        <label className="block text-xs text-gray-500 mb-1">Check In</label>
        <DatePicker
          selected={values.checkIn}
          onChange={date => setValues(prev => ({ ...prev, checkIn: date }))}
          dateFormat="EEE, MMM d"
          className="border rounded-md px-4 py-2 w-full text-gray-700"
          placeholderText="Select date"
        />
      </div>

      {/* Check Out */}
      <div>
        <label className="block text-xs text-gray-500 mb-1">Check Out</label>
        <DatePicker
          selected={values.checkOut}
          onChange={date => setValues(prev => ({ ...prev, checkOut: date }))}
          dateFormat="EEE, MMM d"
          className="border rounded-md px-4 py-2 w-full text-gray-700"
          placeholderText="Select date"
        />
      </div>

      {/* Guests */}
      <div className="relative w-1/4">
        <div
          className="border rounded-md px-4 py-2 cursor-pointer text-gray-700"
          onClick={() =>
            setDropdown(prev => ({ ...prev, guests: !prev.guests }))
          }
        >
          <span className="block text-xs text-gray-500">Guests & Rooms</span>
          {`${totalGuests.adults} Adults, ${totalGuests.children} Children, ${totalRooms} Room${totalRooms > 1 ? "s" : ""}`}
        </div>
        {dropdown.guests && (
          <div className="absolute bg-white border rounded-md mt-1 shadow-lg w-full z-10 p-4">
            {rooms.map((room, index) => (
              <div key={index} className="mb-4">
                <div className="flex justify-between items-center">
                  <h4 className="font-semibold">{`Room ${index + 1}`}</h4>
                  {index > 0 && (
                    <button
                      className="text-red-600 text-sm"
                      onClick={() => removeRoom(index)}
                    >
                      Remove
                    </button>
                  )}
                </div>
                <div className="flex items-center justify-between mt-2">
                  <span>Adults</span>
                  <div className="flex items-center gap-2">
                    <button
                      className="border rounded-full w-6 h-6 flex items-center justify-center"
                      onClick={() => updateRoomCount(index, "adults", "decrement")}
                    >
                      -
                    </button>
                    <span>{room.adults}</span>
                    <button
                      className="border rounded-full w-6 h-6 flex items-center justify-center"
                      onClick={() => updateRoomCount(index, "adults", "increment")}
                    >
                      +
                    </button>
                  </div>
                </div>
                <div className="flex items-center justify-between mt-2">
                  <span>Children</span>
                  <div className="flex items-center gap-2">
                    <button
                      className="border rounded-full w-6 h-6 flex items-center justify-center"
                      onClick={() => updateRoomCount(index, "children", "decrement")}
                    >
                      -
                    </button>
                    <span>{room.children}</span>
                    <button
                      className="border rounded-full w-6 h-6 flex items-center justify-center"
                      onClick={() => updateRoomCount(index, "children", "increment")}
                    >
                      +
                    </button>
                  </div>
                </div>
              </div>
            ))}
            <button
              className="text-blue-600 text-sm underline mb-4"
              onClick={addRoom}
            >
              Add Room
            </button>
            <button
              onClick={() => setDropdown(prev => ({ ...prev, guests: false }))}
              className="bg-purple-600 text-white px-6 py-2 rounded-md w-full"
            >
              Apply
            </button>
          </div>
        )}
      </div>

      {/* Search Button */}
      <button
        onClick={showResult}
        className="bg-purple-600 text-white px-6 py-2 rounded-md flex items-center"
      >
        <span className="mr-2">🔍</span> Search
      </button>
    </div>
  );
};

export default SearchBar;
