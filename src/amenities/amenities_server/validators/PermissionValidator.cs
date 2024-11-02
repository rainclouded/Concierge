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

        public PermissionValidator(PermissionClient permissionClient)
        {
            _permClient = permissionClient;
        }

        public bool ValidatePermissions(string permissionName, string sessionKey)
        {
            var permData = _permClient.GetSessionData(sessionKey).GetAwaiter().GetResult();
            var permissionList = permData?.Data?.SessionData?.SessionPermissionList ?? [];
            Console.WriteLine("Perms: "+string.Join(", ", permissionList));
            return permissionList.Contains(permissionName);
        }
    }
}
