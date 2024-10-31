using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.IdentityModel.Tokens.Jwt;
using Microsoft.IdentityModel.Tokens;
using System.Security.Cryptography;
using Newtonsoft.Json;
using amenities_server.services;
using amenities_server.exceptions;
using amenities_server.model;

namespace amenities_server.validators
{
    public class PermissionValidator : IPermissionValidator
    {
        private readonly PermissionClient _permClient;
        private ECDsa? _publicKey;
        private int permissionPerIndex = 30; //each integer in the permission string contains 30 permissions and is 30 bits long
        Dictionary<string, int> permissionNames;
        private DateTime? _expiration;

        public PermissionValidator(PermissionClient permissionClient) 
        {
            _permClient = permissionClient;
            permissionNames = [];
            _publicKey = null;
            _expiration = null;
        }

        public bool LocalValidatePermissions(string permissionName, string sessionKey)
        {
            var value = false;
            Console.WriteLine("Searching for permissions");
            if (_expiration==null || _expiration < DateTime.Now) 
            {
                var permNameTask = _permClient.GetPermissionList();
                var keyTask = _permClient.GetPublicKey();

                permissionNames = permNameTask.GetAwaiter().GetResult();
                var publicKeyStr = keyTask.GetAwaiter().GetResult();
                _publicKey = ECDsa.Create();
                _publicKey.ImportFromPem(publicKeyStr.AsSpan());
                if (publicKeyStr != "" && permissionNames.Count != 0)
                {
                    _expiration = DateTime.Now.AddSeconds(30);
                }
                else
                {
                    Console.WriteLine("Could not connect to Permissions service");
                    _expiration = DateTime.Now;
                    throw new ConnectionError("Could not connect to Permissions Service");
                }
                Console.WriteLine($"Found permissions: {publicKeyStr}");
            }


            try
            {
                var token = new JwtSecurityToken(sessionKey);
                var tokenHandler = new JwtSecurityTokenHandler();
                var validationParams = new TokenValidationParameters()
                {
                    IssuerSigningKey = new ECDsaSecurityKey(_publicKey!),
                    ValidateAudience = false,
                    ValidateIssuer = false
                };

                SecurityToken sToken;
                var claims = tokenHandler.ValidateToken(sessionKey, validationParams, out sToken);
                
                Console.WriteLine(claims);

                var permissionArrayClaim = claims.Claims.FirstOrDefault(c => c.Type == "permissionString");

                if (permissionArrayClaim != null)
                {

                    int? permissionArray = JsonConvert.DeserializeObject<int>(permissionArrayClaim.Value);
                    Console.WriteLine(permissionArray);
                    if (permissionNames.ContainsKey(permissionName))
                    {
                        Console.WriteLine("PermissionArray contains key");
                        var permId = permissionNames[permissionName];
                        var index = permId / permissionPerIndex;
                        if (permissionArray != null)
                        {
                            value = 0 < (permissionArray & (1 << permId % permissionPerIndex));
                            Console.WriteLine($"Found value: {value}");
                        }
                    }
                }
            } catch (Exception e) { 
                Console.WriteLine(e.ToString());
            } //Return false if the Jwt token does work for any reason
            return value;
        }
    }
}
