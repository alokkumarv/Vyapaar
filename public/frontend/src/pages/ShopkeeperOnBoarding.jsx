import React from 'react'

const ShopkeeperOnBoarding = () => {
    return (
        <form id="registrationForm">
            <h2>Shopkeeper Registration</h2>
            {/* <!-- Personal Info --> */}
            <div class="section">
                <h3>Personal Information</h3>
                <label for="full_name">Full Name</label>
                <input type="text" id="full_name" name="personal_info.full_name" required></input>
                <label for="phone">Phone</label>
                <input type="tel" id="phone" name="personal_info.phone" required></input>
                <label for="email">Email</label>
                <input type="email" id="email" name="personal_info.email" required></input>

                <label for="password">Password</label>
                <input type="password" id="password" name="personal_info.password" required></input>

                <label for="confirm_password">Confirm Password</label>
                <input type="password" id="confirm_password" name="personal_info.confirm_password" required></input>

                {/* <!-- Shop Info --> */}
                <div class="section">
                    <h3>Shop Information</h3>
                    <label for="shop_name">Shop Name</label>
                    <input type="text" id="shop_name" name="shop_info.shop_name" required></input>

                    <label for="shop_category">Shop Category</label>
                    <select id="shop_category" name="shop_info.shop_category" required>
                        <option value="">Select</option>
                        <option value="Grocery">Grocery</option>
                        <option value="Clothing">Clothing</option>
                        <option value="Electronics">Electronics</option>
                        <option value="Pharmacy">Pharmacy</option>
                        <option value="Others">Others</option>
                    </select>

                    <label for="shop_address">Shop Address</label>
                    <textarea id="shop_address" name="shop_info.shop_address" required></textarea>

                    <label for="city">City</label>
                    <input type="text" id="city" name="shop_info.city" required></input>

                    <label for="state">State</label>
                    <input type="text" id="state" name="shop_info.state" required></input>

                    <label for="pincode">Pincode</label>
                    <input type="text" id="pincode" name="shop_info.pincode" required></input>

                    <label for="opening_time">Opening Time</label>
                    <input type="time" id="opening_time" name="shop_info.opening_time"></input>

                    <label for="closing_time">Closing Time</label>
                    <input type="time" id="closing_time" name="shop_info.closing_time"></input>

                    <label for="license_number">License Number</label>
                    <input type="text" id="license_number" name="shop_info.license_number"></input>
                </div>

                {/* <!-- Payment Info --> */}
                <div class="section">
                    <h3>Payment Information (Optional)</h3>
                    <label for="account_holder">Account Holder Name</label>
                    <input type="text" id="account_holder" name="payment_info.account_holder"></input>

                    <label for="account_number">Account Number</label>
                    <input type="text" id="account_number" name="payment_info.account_number"></input>

                    <label for="ifsc_code">IFSC Code</label>
                    <input type="text" id="ifsc_code" name="payment_info.ifsc_code"></input>

                    <label for="upi_id">UPI ID</label>
                    <input type="text" id="upi_id" name="payment_info.upi_id"></input>
                </div>
                {/* <!-- Document Info --> */}
                <div class="section">
                    <h3>Document Uploads</h3>
                    <label for="shop_photo_url">Shop Photo URL</label>
                    <input type="text" id="shop_photo_url" name="document_info.shop_photo_url"></input>

                    <label for="owner_id_proof_url">Owner ID Proof URL</label>
                    <input type="text" id="owner_id_proof_url" name="document_info.owner_id_proof_url"></input>

                    <label for="registration_cert_url">Shop Registration Cert URL</label>
                    <input type="text" id="registration_cert_url" name="document_info.registration_cert_url"></input>
                </div>

                {/* <!-- Location Info --> */}
                <div class="section">
                    <h3>Location Information</h3>
                    <label for="latitude">Latitude</label>
                    <input type="number" step="any" id="latitude" name="location_info.latitude"></input>

                    <label for="longitude">Longitude</label>
                    <input type="number" step="any" id="longitude" name="location_info.longitude"></input>
                </div>

                <button type="submit">Register</button>


            </div>

        </form>
    )
}

export default ShopkeeperOnBoarding