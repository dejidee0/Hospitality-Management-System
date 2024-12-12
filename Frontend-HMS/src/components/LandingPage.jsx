import React, { useState } from 'react';

const HotelLandingPage = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    phone: '',
    message: ''
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Form submitted:', formData);
    alert('Thank you for your inquiry! We will contact you soon.');
  };

  return (
    <div className="font-sans text-gray-900">
      {/* Sticky Navigation */}
      <nav className="fixed top-0 left-0 right-0 bg-white shadow-md z-50">
        <div className="max-w-6xl mx-auto px-4 py-3 flex justify-between items-center">
          <div className="text-2xl font-bold">LuxeStay Hotel</div>
          <div className="space-x-6">
            {['Rooms', 'Amenities', 'Gallery', 'Contact'].map(item => (
              <a 
                key={item} 
                href={item.toLowerCase()} 
                className="hover:text-blue-600 transition"
              >
                {item}
              </a>
            ))}
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <header 
        id="home" 
        className="relative h-screen flex items-center justify-center text-white"
        style={{
          backgroundImage: `url('/api/placeholder/1920/1080')`,
          backgroundSize: 'cover',
          backgroundPosition: 'center',
          marginTop: '-64px'
        }}
      >
        <div className="absolute inset-0 bg-black opacity-50"></div>
        <div className="relative z-10 text-center max-w-2xl px-4">
          <h1 className="text-5xl font-bold mb-4">Escape to Luxury</h1>
          <p className="text-xl mb-8">Experience unparalleled comfort and elegance at LuxeStay Hotel</p>
          <a 
            href="#contact" 
            className="bg-blue-600 text-white px-8 py-3 rounded-full text-lg hover:bg-blue-700 transition"
          >
            Book Now
          </a>
        </div>
      </header>

      {/* Key Benefits */}
      <section id="amenities" className="py-16 bg-gray-100">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className="text-4xl text-center mb-12 font-semibold">Why Choose LuxeStay</h2>
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { 
                title: 'Premium Rooms', 
                description: 'Luxurious rooms with stunning views and modern amenities',
                image: '/api/placeholder/400/300?1'
              },
              { 
                title: 'Gourmet Dining', 
                description: 'World-class restaurants offering exquisite culinary experiences',
                image: `/api/placeholder/400/300?2`
              },
              { 
                title: 'Wellness Center', 
                description: 'State-of-the-art spa and fitness facilities',
                image: `/api/placeholder/400/300?3`
              }
            ].map((benefit, index) => (
              <div 
                key={index} 
                className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition"
              >
                <img 
                  src={benefit.image} 
                  alt={benefit.title} 
                  className="w-full h-48 object-cover"
                />
                <div className="p-6">
                  <h3 className="text-2xl font-bold mb-4">{benefit.title}</h3>
                  <p className="text-gray-600">{benefit.description}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Gallery */}
      <section id="gallery" className="py-16">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className="text-4xl text-center mb-12 font-semibold">Our Spaces</h2>
          <div className="grid md:grid-cols-4 gap-4">
            {[4, 5, 6, 7, 8, 9, 10, 11].map((num) => (
              <img 
                key={num} 
                src= '/api/placeholder/300/200?${num}'
                alt='Hotel Space ${num}'
                className="w-full h-48 object-cover rounded-lg hover:scale-105 transition"
              />
            ))}
          </div>
        </div>
      </section>

      {/* Testimonials */}
      <section className="py-16 bg-gray-100">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className="text-4xl text-center mb-12 font-semibold">Guest Experiences</h2>
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { 
                name: 'Emily Roberts', 
                quote: 'The most incredible hotel experience I\'ve ever had. Absolutely stunning!',
                image: '/api/placeholder/200/200?12'
              },
              { 
                name: 'Michael Chen', 
                quote: 'Perfect location, incredible service, and luxurious rooms.',
                image: '/api/placeholder/200/200?13'
              },
              { 
                name: 'Sarah Thompson', 
                quote: 'A true gem. Every detail was perfect and the staff was exceptional.',
                image: '/api/placeholder/200/200?14'
              }
            ].map((review, index) => (
              <div 
                key={index} 
                className="bg-white p-6 rounded-lg shadow-md text-center"
              >
                <img 
                  src={review.image} 
                  alt={review.name}
                  className="w-24 h-24 rounded-full mx-auto mb-4 object-cover"
                />
                <blockquote className="italic mb-4">"{review.quote}"</blockquote>
                <p className="font-bold">- {review.name}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Packages */}
      <section id="packages" className="py-16">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className="text-4xl text-center mb-12 font-semibold">Our Packages</h2>
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { 
                title: 'Weekend Escape', 
                price: '$299',
                features: ['2 nights stay', 'Breakfast included', 'Spa access']
              },
              { 
                title: 'Romantic Getaway', 
                price: '$499',
                features: ['3 nights stay', 'Couples massage', 'Gourmet dinner']
              },
              { 
                title: 'Luxury Suite', 
                price: '$799',
                features: ['4 nights stay', 'Ocean view', 'Personal concierge']
              }
            ].map((pkg, index) => (
              <div 
                key={index} 
                className="bg-white border rounded-lg p-6 text-center hover:shadow-xl transition"
              >
                <h3 className="text-2xl font-bold mb-4">{pkg.title}</h3>
                <p className="text-4xl font-bold text-blue-600 mb-6">{pkg.price}</p>
                <ul className="mb-6 space-y-2">
                  {pkg.features.map((feature, i) => (
                    <li key={i} className="text-gray-600">{feature}</li>
                  ))}
                </ul>
                <a 
                  href="#contact" 
                  className="w-full block bg-blue-600 text-white py-3 rounded-full hover:bg-blue-700 transition"
                >
                  Book Now
                </a>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Contact Form */}
      <section id="contact" className="py-16 bg-gray-100">
        <div className="max-w-xl mx-auto px-4">
          <h2 className="text-4xl text-center mb-12 font-semibold">Contact Us</h2>
          <form 
            onSubmit={handleSubmit} 
            className="bg-white p-8 rounded-lg shadow-md"
          >
            <div className="mb-4">
              <label htmlFor="name" className="block mb-2">Name</label>
              <input 
                type="text" 
                id="name" 
                name="name"
                value={formData.name}
                onChange={handleInputChange}
                required 
                className="w-full px-3 py-2 border rounded-md"
                placeholder="Your Name"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="email" className="block mb-2">Email</label>
              <input 
                type="email" 
                id="email" 
                name="email"
                value={formData.email}
                onChange={handleInputChange}
                required 
                className="w-full px-3 py-2 border rounded-md"
                placeholder="your.email@example.com"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="phone" className="block mb-2">Phone</label>
              <input 
                type="tel" 
                id="phone" 
                name="phone"
                value={formData.phone}
                onChange={handleInputChange}
                className="w-full px-3 py-2 border rounded-md"
                placeholder="(Optional) Your Phone Number"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="message" className="block mb-2">Message</label>
              <textarea 
                id="message" 
                name="message"
                value={formData.message}
                onChange={handleInputChange}
                className="w-full px-3 py-2 border rounded-md"
                placeholder="Your inquiry or special requests"
                rows="4"
              ></textarea>
            </div>
            <button 
              type="submit" 
              className="w-full bg-blue-600 text-white py-3 rounded-full hover:bg-blue-700 transition"
            >
              Send Inquiry
            </button>
          </form>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-white py-12">
        <div className="max-w-6xl mx-auto px-4 grid md:grid-cols-3 gap-8">
          <div>
            <h4 className="text-xl font-bold mb-4">LuxeStay Hotel</h4>
            <p>123 Luxury Lane, Elegant City</p>
            <p>Phone: +1 (555) 123-4567</p>
            <p>Email: reservations@luxestay.com</p>
          </div>
          <div>
            <h4 className="text-xl font-bold mb-4">Quick Links</h4>
            <ul className="space-y-2">
              <li><a href="#" className="hover:text-blue-400">Privacy Policy</a></li>
              <li><a href="#" className="hover:text-blue-400">Terms of Service</a></li>
              <li><a href="#" className="hover:text-blue-400">FAQ</a></li>
            </ul>
          </div>
          <div>
            <h4 className="text-xl font-bold mb-4">Follow Us</h4>
            <div className="flex space-x-4">
              {['Facebook', 'Instagram', 'Twitter'].map((platform) => (
                <a 
                  key={platform} 
                  href="#" 
                  className="hover:text-blue-400"
                >
                  {platform}
                </a>
              ))}
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default HotelLandingPage;