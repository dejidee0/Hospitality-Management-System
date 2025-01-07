import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";
import Navbar from "../auth/components/Navbar";
import hotels from "../../assets/imgs/hotels.png"
import Loading from "../../components/ButtonLoader";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { CgDanger } from "react-icons/cg";
import googleIcon from "../../assets/google_icon.svg";
import facebookIcon from "../../assets/facebook_icon.svg";


const Home = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
  
    const { isLoading, error } = useSelector((state) => state.auth);
  
    const handleSubmit = async (e) => {
      e.preventDefault();
  
      if (validate()) {
        dispatch(signup(formData))
          .unwrap()
          .then(() => {
            navigate("/signin");
          })
          .catch((err) => {
            console.error("Signup error:", err);
          });
      }
    };
  
    return (
      <div>
      <Navbar />
    
      <section class="relative w-full">
        <div class="relative">
          <img
            class="w-full h-[500px] object-cover"
            src={hotels}
            alt="Hero Image"
          />

          <div class="absolute inset-0 flex flex-col justify-center items-center text-white bg-black/50">
            <h1 class="text-4xl font-bold">A piece of paradise just for you</h1>
          </div>

          <div className="absolute mt-[400px] inset-0 flex flex-col justify-center items-center">
            <div className="w-full bg-white">
              <ul className="flex items-center justify-center">
                <li className="w-5 p-5">Hotels</li>
                <li className="w-5 p-5">Flights</li>
                <li className="w-5 p-5">Rentals</li>
                <li className="w-5 p-5">Events</li>
              </ul>
            </div>
            <div class=" bg-white p-6 rounded-lg shadow-md w-full max-w-4xl flex flex-col md:flex-row gap-4">
              <div class="flex-grow">
                <label for="destination" class="block text-sm font-medium text-gray-700">
                  Where To
                </label>
                <input
                  id="destination"
                  type="text"
                  placeholder="Search Destination"
                  class="w-full p-3 border border-gray-300 rounded-md"
                />
              </div>
              <div>
                <label for="checkin" class="block text-sm font-medium text-gray-700">
                  Check In
                </label>
                <input
                  id="checkin"
                  type="date"
                  class="w-full p-3 border border-gray-300 rounded-md"
                />
              </div>
              <div>
                <label for="checkout" class="block text-sm font-medium text-gray-700">
                  Check Out
                </label>
                <input
                  id="checkout"
                  type="date"
                  class="w-full p-3 border border-gray-300 rounded-md"
                />
              </div>
              <div>
                <label for="guests" class="block text-sm font-medium text-gray-700">
                  Guests & Rooms
                </label>
                <select
                  id="guests"
                  class="w-full p-3 border border-gray-300 rounded-md"
                >
                  <option>2 Adults, 1 Room</option>
                  <option>1 Adult, 1 Room</option>
                  <option>3 Adults, 2 Rooms</option>
                </select>
              </div>
              <button class="bg-purple-600 hover:bg-purple-700 text-white font-bold py-3 px-6 rounded-lg">
                Search
              </button>
            </div>
          </div>
          
        </div>
      </section>
{/* 
      <section class="bg-gray-50 relative -mt-16 z-10 rounded-t-lg">
        <div class="py-10 flex justify-center space-x-8">

          <div class="text-center">
            <img
              src="path-to-icon1"
              alt="Icon"
              class="h-12 w-12 mx-auto"
            />
            <h3 class="text-lg font-bold mt-4">Seamless Booking Experience</h3>
            <p class="text-gray-600 mt-2">
              Our easy-to-use platform guarantees a smooth and secure booking process.
            </p>
          </div>
          <div class="text-center">
            <img
              src="path-to-icon2"
              alt="Icon"
              class="h-12 w-12 mx-auto"
            />
            <h3 class="text-lg font-bold mt-4">Secure & Simple Payments</h3>
            <p class="text-gray-600 mt-2">
              Easily manage bookings with secure payments and instant updates.
            </p>
          </div>
          <div class="text-center">
            <img
              src="path-to-icon3"
              alt="Icon"
              class="h-12 w-12 mx-auto"
            />
            <h3 class="text-lg font-bold mt-4">Customer Support</h3>
            <p class="text-gray-600 mt-2">
              Assistance with bookings, payments, and questions—always here to help.
            </p>
          </div>
        </div>
      </section> */}
    </div>
    
     
     
    );
  };

export default Home;
