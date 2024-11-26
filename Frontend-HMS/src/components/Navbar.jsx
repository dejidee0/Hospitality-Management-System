import { useState } from "react";
import logoImg from "../assets/Findpeace-logo.svg";
import { FaChevronDown } from "react-icons/fa";
import Flag from "react-world-flags"; // For displaying flags

const Navbar = () => {
  const [country, setCountry] = useState("NG"); // Default country (Nigeria)
  const [language, setLanguage] = useState("EN"); // Default language
  const [currency, setCurrency] = useState("NGN"); // Default currency
  const [openDropdown, setOpenDropdown] = useState(""); // Tracks which dropdown is open

  // Dummy data for dropdowns
  const countries = [
    { name: "Nigeria", code: "NG" },
    { name: "Benin Republic", code: "BJ" },
  ];
  const languages = ["EN", "FR"];
  const currencies = ["NGN", "CFA"];

  // Handle dropdown toggle
  const toggleDropdown = (dropdown) => {
    setOpenDropdown(openDropdown === dropdown ? "" : dropdown); // Close if open, open if closed
  };

  // Close dropdown when clicking outside
  const handleBlur = () => {
    setTimeout(() => setOpenDropdown(""), 150); // Timeout ensures dropdown stays open during a click
  };

  return (
    <header
      style={{
        position: "fixed",
        width: "100vw",
        backgroundColor: "#ffffff",
        height: "80px",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        padding: "0px 6.6rem",
        zIndex: 1000,
        boxShadow: "0px 2px 5px rgba(0, 0, 0, 0.1)",
      }}
    >
      {/* Logo */}
      <div>
        <img src={logoImg} alt="Find peace logo" />
      </div>

      {/* Flag, Language, and Currency Dropdowns */}
      <div
        style={{
          display: "flex",
          alignItems: "center",
          gap: "17px",
          width: "255px",
        }}
      >
        {/* Country Dropdown */}
        <div style={{ position: "relative" }} onBlur={handleBlur} tabIndex={0}>
          <div
            style={{
              display: "flex",
              alignItems: "center",
              gap: "8px",
              cursor: "pointer",
            }}
            onClick={() => toggleDropdown("country")}
          >
            <Flag code={country} style={{ width: "24px", height: "16px" }} />
            <FaChevronDown size={12} />
          </div>
          {/* Dropdown Menu */}
          {openDropdown === "country" && (
            <div
              style={{
                position: "absolute",
                top: "100%",
                left: 0,
                backgroundColor: "#ffffff",
                boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
                borderRadius: "4px",
                marginTop: "8px",
                animation: "fadeIn 0.2s ease-in-out",
                zIndex: 100,
              }}
            >
              {countries.map((c) => (
                <div
                  key={c.code}
                  style={{
                    padding: "8px 16px",
                    cursor: "pointer",
                    display: "flex",
                    alignItems: "center",
                    gap: "8px",
                  }}
                  onClick={() => {
                    setCountry(c.code);
                    setOpenDropdown("");
                  }}
                >
                  <Flag
                    code={c.code}
                    style={{ width: "24px", height: "16px" }}
                  />
                  <span
                    style={{
                      fontSize: "14px",
                      fontWeight: "500",
                      lineHeight: "19.6px",
                    }}
                  >
                    {c.name}
                  </span>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Language Dropdown */}
        <div style={{ position: "relative" }} onBlur={handleBlur} tabIndex={0}>
          <div
            style={{
              display: "flex",
              alignItems: "center",
              gap: "8px",
              cursor: "pointer",
              borderLeft: "1px solid #b2b3b3",
              borderRight: "1px solid # b2b3b3",
              paddingLeft: "8px",
              paddingRight: "10px",
            }}
            onClick={() => toggleDropdown("language")}
          >
            <span
              style={{
                fontSize: "14px",
                fontWeight: "500",
                lineHeight: "19.6px",
              }}
            >
              {language}
            </span>
            <FaChevronDown size={12} />
          </div>
          {/* Dropdown Menu */}
          {openDropdown === "language" && (
            <div
              style={{
                position: "absolute",
                top: "100%",
                left: 0,
                backgroundColor: "#ffffff",
                boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
                borderRadius: "4px",
                marginTop: "8px",
                animation: "fadeIn 0.2s ease-in-out",
                zIndex: 100,
              }}
            >
              {languages.map((lang) => (
                <div
                  key={lang}
                  style={{
                    padding: "8px 16px",
                    cursor: "pointer",
                  }}
                  onClick={() => {
                    setLanguage(lang);
                    setOpenDropdown("");
                  }}
                >
                  {lang}
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Currency Dropdown */}
        <div style={{ position: "relative" }} onBlur={handleBlur} tabIndex={0}>
          <div
            style={{
              display: "flex",
              alignItems: "center",
              gap: "8px",
              cursor: "pointer",
            }}
            onClick={() => toggleDropdown("currency")}
          >
            <span
              style={{
                fontSize: "14px",
                fontWeight: "500",
                lineHeight: "19.6px",
              }}
            >
              {currency}
            </span>
            <FaChevronDown size={12} />
          </div>
          {/* Dropdown Menu */}
          {openDropdown === "currency" && (
            <div
              style={{
                position: "absolute",
                top: "100%",
                left: 0,
                backgroundColor: "#ffffff",
                boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
                borderRadius: "4px",
                marginTop: "8px",
                animation: "fadeIn 0.2s ease-in-out",
                zIndex: 100,
              }}
            >
              {currencies.map((cur) => (
                <div
                  key={cur}
                  style={{
                    padding: "8px 16px",
                    cursor: "pointer",
                  }}
                  onClick={() => {
                    setCurrency(cur);
                    setOpenDropdown("");
                  }}
                >
                  {cur}
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

export default Navbar;
